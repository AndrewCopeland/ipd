package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AndrewCopeland/go-ipdetective"
)

var (
	jsonFlag        = flag.Bool("json", false, "returns ip information in json format")
	csvFlag         = flag.Bool("csv", false, "returns ip information in csv format")
	helpFlag        = flag.Bool("help", false, "shows this help menu")
	helpDescription = `ipd - a ip bot detection and ip geolocation command line utility

set api key by using the environment variable 'IPDETECTIVE_API_KEY'

get your FREE api key at https://ipdetective.io

Example usage:
  ipd 			# get my ip info	
  ipd 8.8.8.8 		# get ip info about 8.8.8.8
  ipd -csv 8.8.8.8  	# get ip info in csv format
  ipd -json 8.8.8.8  	# get ip info in json format
  
  # create CSV file of all unique vistors from nginx logs
  cat /var/log/nginx/access.log | awk '{print $1}' | sort | uniq | ipd -csv > unique_vistors.csv
`
)

func main() {
	ip := ""
	flag.Parse()
	if helpFlag != nil && *helpFlag {
		fmt.Println(helpDescription)
		return
	}
	apiKey := os.Getenv("IPDETECTIVE_API_KEY")
	if len(flag.Args()) >= 1 {
		ip = flag.Arg(0)
	}

	cfg := ipdetective.NewConfiguration()
	if apiKey != "" {
		cfg.AddDefaultHeader("x-api-key", apiKey)
	}
	client := ipdetective.NewAPIClient(cfg).DefaultAPI
	ctx := context.Background()

	if ok := processStdIn(ctx, client); ok {
		return
	}
	if ip == "" {
		processMyIP(ctx, client)
		return
	}
	processIP(ctx, client, ip)
}

func processStdIn(ctx context.Context, client *ipdetective.DefaultAPIService) bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("could not read from stdin. %s", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			processIP(ctx, client, line)
		}
		return true
	}

	return false
}

func processMyIP(ctx context.Context, client *ipdetective.DefaultAPIService) {
	resp, httpResp, err := client.GetMyIP(ctx).Info(true).Execute()
	if err != nil {
		log.Fatalf("failed to get my ip address. %s", err)
	}
	if httpResp.StatusCode != 200 {
		log.Fatalf("invalid status code was returned. %d", httpResp.StatusCode)
	}

	printIPResponse(resp)
}

func processIP(ctx context.Context, client *ipdetective.DefaultAPIService, ip string) {
	resp, httpResp, err := client.GetIP(ctx, ip).Info(true).Execute()
	if err != nil {
		log.Fatalf("failed to get info about ip address. %s", err)
	}
	if httpResp.StatusCode != 200 {
		log.Fatalf("invalid status code was returned. %d", httpResp.StatusCode)
	}
	printIPResponse(resp)
}

func printIPResponse(resp *ipdetective.IPResponse) {
	if csvFlag != nil && *csvFlag {
		fmt.Printf("\"%s\",\"%t\",\"%s\",\"%d\",\"%s\",\"%s\",\"%s\"\n",
			resp.Ip, resp.Bot, fromPtr(resp.Type), fromPtr(resp.Asn), fromPtr(resp.AsnDescription),
			fromPtr(resp.CountryCode), fromPtr(resp.CountryName))
		return
	}
	if jsonFlag != nil && *jsonFlag {
		resp, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Fatalf("could not marshal ipdetective response. %s", err)
		}
		fmt.Println(string(resp))
		return
	}

	// human readable format
	fmt.Printf("IP:\t\t%s\n", resp.Ip)
	fmt.Printf("Bot:\t\t%t\n", resp.Bot)
	fmt.Printf("Type:\t\t%s\n", fromPtr(resp.Type))
	fmt.Printf("ASN:\t\t%d\n", fromPtr(resp.Asn))
	fmt.Printf("ASN Desc:\t%s\n", fromPtr(resp.AsnDescription))
	fmt.Printf("Country Code:\t%s\n", fromPtr(resp.CountryCode))
	fmt.Printf("Country Name:\t%s\n", fromPtr(resp.CountryName))
}

func fromPtr[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

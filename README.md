# ipd

ipd - ip bot detection and ip geolocation command line utility

set api key by using the environment variable 'IPDETECTIVE_API_KEY'

get your free api key at https://ipdetective.io

```
Example usage:
	ipd                 # get my ip info
	ipd 8.8.8.8 		# get ip info about 8.8.8.8

	ipd -csv 8.8.8.8  	# get ip info in csv format
	ipd -json 8.8.8.8  	# get ip info in json format

	# create CSV file of all unique vistors from nginx logs
	cat /var/log/nginx/access.log | awk '{print $1}' | sort | uniq | ipd -csv > unique_vistors.csv
```

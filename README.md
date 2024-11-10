# ipd

ipd - ip bot detection and ip geolocation command line utility

If you hit limit restriction set your environment variable 'IPDETECTIVE_API_KEY'

Get your free api key at https://ipdetective.io

## Installation
On linux and mac perform the following command:
```
curl -s "https://raw.githubusercontent.com/AndrewCopeland/ipd/refs/heads/main/install.sh" | bash
```

To install on windows you must download the executable file from the [archive](https://github.com/AndrewCopeland/ipd/releases)


## Usage
Get my current machines IP address:
```bash
ipd
```

Get other machines IP address:
```bash
ipd 8.8.8.8
```

Get other machines IP address in JSON format:
```bash
ipd -json 8.8.8.8
```

Get other machines IP address in CSV format:
```bash
ipd -csv 8.8.8.8
```

Get all unique nginx vistors and output to CSV
```bash
cat /var/log/nginx/access.log | awk '{print $1}' | sort | uniq | ipd -csv > unique_vistors.csv
```



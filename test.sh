#!/bin/bash
set -e

TEST_GOOGLE_SUCCESS=$(cat <<-END
IP:		8.8.8.8
Bot:		true
Type:		datacenter
ASN:		15169
ASN Desc:	GOOGLE
Country Code:	US
Country Name:	United States
END
)


TEST_GOOGLE_CSV_SUCCESS=$(cat <<-END
"8.8.8.8","true","datacenter","15169","GOOGLE","US","United States"
END
)

TEST_GOOGLE_JSON_SUCCESS=$(cat <<-END
{
  "asn": 15169,
  "asn_description": "GOOGLE",
  "bot": true,
  "country_code": "US",
  "country_name": "United States",
  "ip": "8.8.8.8",
  "type": "datacenter"
}
END
)

TEST_CSV_PIPE_SUCCESS=$(cat <<-END
"8.8.8.8","true","datacenter","15169","GOOGLE","US","United States"
"1.1.1.1","true","datacenter","13335","CLOUDFLARENET","AU","Australia"
END
)

assert () {
    name="$1"
    actual="$2"
    expected="$3"
    if [[ "$actual" != "$expected" ]]; then
        echo "FAILED - $name - actual is not expected"
        echo "  actual: $actual"
        echo "  expected: $expected"
        return 1
    fi
    echo "PASSED - $name"
    return 0
}

assert "ip_success" "$(./ipd 8.8.8.8)" "$TEST_GOOGLE_SUCCESS"
assert "ip_csv_success" "$(./ipd -csv 8.8.8.8)" "$TEST_GOOGLE_CSV_SUCCESS"
assert "ip_json_success" "$(./ipd -json 8.8.8.8)" "$TEST_GOOGLE_JSON_SUCCESS"

# test out piping and returning CSV
rm -f test.csv
rm -f test-out.csv
echo "8.8.8.8" >> test.csv
echo "1.1.1.1" >> test.csv
cat test.csv | ./ipd -csv > test-out.csv
assert "ip_csv_pipe_success" "$(cat test-out.csv)" "$TEST_CSV_PIPE_SUCCESS"
rm -f test.csv
rm -f test-out.csv

# set invalid api key and expect exit code of 1
set +e
export IPDETECTIVE_API_KEY="this-is-not-real"
./ipd 8.8.8.8 &> /dev/null
assert "invalid_api_key" "$(echo $?)" "1"
unset IPDETECTIVE_API_KEY
set -e
#!/bin/bash

os=$(uname)
architecture=$(uname -m)

fileName="ipd_${os}_${architecture}.tar.gz"
url="https://github.com/AndrewCopeland/ipd/releases/download/v0.0.3/$fileName"

echo "downloading ipd from $url"
curl -O "$url"
tar -xvzf "$fileName"
sudo mv "./ipd" "/usr/local/bin/ipd"
rm "$fileName"
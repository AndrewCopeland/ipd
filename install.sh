#!/bin/bash

os=$(uname)
architecture=$(uname -m)

fileName="ipd_${os}_${architecture}.tar.gz"
url="https://github.com/AndrewCopeland/ipd/releases/download/v0.0.4/$fileName"

echo "downloading ipd from $url"
curl -s -OL "$url"
echo "ipd tarball has been download"
echo "extracting ipd tarball"
tar -xvzf "$fileName"
echo "ipd tarball extracted"
echo "moving file to /usr/local/bin/ipd"
sudo mv "./ipd" "/usr/local/bin/ipd"
rm "$fileName"
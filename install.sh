#!/bin/bash

os=$(uname)
architecture=$(uname -m)

fileName="ipd_${os}_${architecture}.tar.gz"
url="https://raw.githubusercontent.com/AndrewCopeland/ipd/refs/heads/main/dist/$fileName"

echo "downloading ipd from $url"


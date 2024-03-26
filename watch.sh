#!/bin/sh
file="$1"

echo "The file to run: $file"

# kill old process and start new one
pkill $file
go run "$file"

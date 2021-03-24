#!/bin/bash

version=$(grep "const VERSION" Main.go|awk -F'"' '{print $2}')
for platform in $(go tool dist list)
do
  rm -rf portscan
  goos=$(echo $platform|awk -F'/' '{print $1}')
  goarch=$(echo $platform|awk -F'/' '{print $2}')
  if [ "$goos" == "android" ];then
    continue
  fi
  if [ "$goos" == "js" ];then
    continue
  fi
  if [ "$goos" == "ios" ];then
    continue
  fi

  CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -o portscan Main.go
  echo "tar compress portscan_${version}_${goos}_${goarch}.tar.gz"
  tar -czvf "portscan_${version}_${goos}_${goarch}.tar.gz" portscan README.md
done
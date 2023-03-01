go build -ldflags="-s -w"  -o gapi

if command -v upx > /dev/null 2>&1; then
  WITH_UPX=1
else
  WITH_UPX=0
fi

if [ $WITH_UPX -eq 1 ];then
  upx -9 gapi
fi

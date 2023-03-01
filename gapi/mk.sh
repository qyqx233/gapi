
go build -ldflags="-s -w" *.go -o gapi

if command -v upx > /dev/null 2>&1; then
  WITH_UPX=1
else
  WITH_UPX=1
fi

if [ $WITH_UPX -eq 1 ];then
  upx -9 gapi
fi
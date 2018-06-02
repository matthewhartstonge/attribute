cd cmd\attribute
go build -ldflags="-s -w"
upx --brute attribute.exe

cd ..\..\
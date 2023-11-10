# can use this to build if you don't want to use Makefile

env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hd-wallet wallet/main.go

echo "Done."
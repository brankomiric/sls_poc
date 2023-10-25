# can use this to build if you don't want to use Makefile

env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ipAddresser ip-addresser/main.go
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/geoLocator geo-locator/main.go

echo "Done."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o gohs_min.exe


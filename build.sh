rm bin/forwarder
GOOS=linux go build -o bin/forwarder cmd/forwarder.go
docker build -t kieper/forwarder .
# Forwarder

Simple project for forwarding requests

By default start on `8080` port, can be overridden by setting `PORT` variable.
To forwards you need to provide `path` query parameter, sample request:
```bash
curl http://localhost:8080?path=http://google.pl
```

## To run on localhost
```bash
go run cmd/forwarder.go
```

## To build docker image run 
```bash
./build.sh
```

## To run docker image

```bash
docker run -p 8080:8080 kieper/forwarder
```
# About

This is a simple example showing [gRPC](https://grpc.io/) [golang](https://golang.org/) service running on [Cloud Run](https://cloud.google.com/run/).

# Requirements

1.  [gcloud](https://cloud.google.com/sdk/gcloud/) with local authentication and GCP billing setup.
2.  [docker](https://www.docker.com/)

# Optional

The following is needed for building your own grpc but already done for this repo example.

Golang protoc commandline tool: [protoc](https://grpc.io/docs/quickstart/go/#install-grpc)

# Setup and usage

## Build Docker image 

```
$ make build
```

## Push Docker image

```
$ make push
```

## Deploy Cloud Run service

```
$ make deploy
```

## Test using local client

```
$ go run client/client.go $(make print-host) hi
```

Expected response where 'hi' is repeated twice to prove processing by the server:

```
2019/12/01 14:19:04 payload:"hihi" 
```


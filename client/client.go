package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"regexp"
	echo "temp/proto"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: client.go HOST PAYLOAD")
		os.Exit(0)
	}
	host := args[0]
	r := regexp.MustCompile("^https://")
	host = r.ReplaceAllString(host, "")
	address := fmt.Sprintf("%s:443", host)
	payload := args[1]
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	client := echo.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Get(ctx, &echo.GetRequest{
		Payload: payload,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
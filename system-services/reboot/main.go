// go run main.go -username ansible -password ansible -target 10.20.30.37:6030
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/golang/glog"
	system "github.com/openconfig/gnoi/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func checkflags(flag ...string) {
	for _, f := range flag {
		if f == "" {
			fmt.Printf("You have an empty flag please fix.")
			os.Exit(1)
		}
	}
}

func main() {
	username := flag.String("username", "", "username for connection to gNOI")
	password := flag.String("password", "", "password for connection to gNOI")
	target := flag.String("target", "", "Target ip or hostname of the device running gNOI")
	flag.Parse()
	conn, err := grpc.Dial(*target, grpc.WithInsecure())
	if err != nil {
		log.Exitf("Failed to %s Error: %v", target, err)
	}
	defer conn.Close()

	// Create the new grpc service connection
	Sys := system.NewSystemClient(conn)
	// pass in context blank information with the timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// cancel when the function is over.
	defer cancel()
	// Since Metadata needs a map to pass into the header of gRPC request create a map for it.
	metamap := make(map[string]string)
	// Set the username and password
	metamap["username"] = *username
	metamap["password"] = *password
	// Set the metadata needed in the metadata package
	md := metadata.New(metamap)
	// set the ctx to use the metadata in every update.
	ctx = metadata.NewOutgoingContext(ctx, md)
	response, err := Sys.Reboot(ctx, &system.RebootRequest{Method: 1, Delay: 1, Force: true})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.String())

}

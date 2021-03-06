package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

//Run - runs the server
func Run(ctx context.Context, port string) struct{} {

	//Server to build up
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Stopped the server: %v", err)
	}
	fmt.Println("Running on port " + port)

	//Passing grpc to the variables
	s := &Server{}
	grpcServer := grpc.NewServer()
	publish.RegisterPublishServiceServer(grpcServer, s)
	grpcServer.Serve(li)

	return <-ctx.Done()
}

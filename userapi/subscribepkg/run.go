package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

//Run - runs the server in the port specified
func Run(ctx context.Context, port string) struct{} {

	//The server to get up
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Server stopped: %v", err)
	}

	//Passing the server to grpc
	s := &Server{}
	grpcServer := grpc.NewServer()
	subscribe.RegisterSubscribeServiceServer(grpcServer, s)
	grpcServer.Serve(li)

	fmt.Printf("Server up on port: %v\n", err)
	return <-ctx.Done()
}

package userfeedpkg

import (
	"fmt"
	userfeed "github.com/WEGOAPP/userapi/protos/userfeed"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

//Run - runs the server on the port
func Run(ctx context.Context, port string) struct{} {

	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Couldnt start server: %v", err)
	}
	s := &Server{}
	grpcServer := grpc.NewServer()
	userfeed.RegisterUserFeedServer(grpcServer, s)
	fmt.Printf("Server up on port %v\n", port)
	grpcServer.Serve(li)

	return <-ctx.Done()
}

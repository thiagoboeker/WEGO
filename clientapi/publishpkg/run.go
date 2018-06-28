package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	grpc "google.golang.org/grpc"
	"net"
)

//Run - runs the server
func Run(port string) error {
	li, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Error running server")
	}
	fmt.Println("Running on port " + port)
	errch := make(chan error)
	s := &Server{}
	grpcServer := grpc.NewServer()
	publish.RegisterPublishServiceServer(grpcServer, s)
	grpcServer.Serve(li)
	return <-errch
}

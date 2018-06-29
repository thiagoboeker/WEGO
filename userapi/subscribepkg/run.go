package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
	grpc "google.golang.org/grpc"
	"net"
)

//Run - runs the server in the port specified
func Run(port string) error {
	li, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Couldnt create server: %v", err)
	}
	errorch := make(chan error)
	s := &Server{}
	grpcServer := grpc.NewServer()
	subscribe.RegisterSubscribeServiceServer(grpcServer, s)
	grpcServer.Serve(li)
	return <-errorch
}

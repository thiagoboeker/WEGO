package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
	"golang.org/x/net/context"
)

//Server - interface handler
type Server struct{}

//SubscribeRequest - grpc function implementation
func (s *Server) SubscribeRequest(ctx context.Context, subblock *subscribe.SubscribeBlock) (*subscribe.Done, error) {
	fmt.Println(subblock)
	return &subscribe.Done{Status: 200}, nil
}

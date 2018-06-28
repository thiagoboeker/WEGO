package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	context "golang.org/x/net/context"
)

//Server - grpc handler
type Server struct{}

//PublishRequest - Publish the post
func (s *Server) PublishRequest(ctx context.Context, pub *publish.PublishBlock) (*publish.Done, error) {
	fmt.Println(pub)
	return &publish.Done{Status: 200}, nil
}

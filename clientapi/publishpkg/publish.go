package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	redis "github.com/garyburd/redigo/redis"
	context "golang.org/x/net/context"
	"strings"
)

//Server - grpc handler
type Server struct{}

//PublishRequest - Publish the post
func (s *Server) PublishRequest(ctx context.Context, pub *publish.PublishBlock) (*publish.Done, error) {
	conn, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		return nil, fmt.Errorf("Error reaching db: %v", err)
	}
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("ZADD", strings.Join([]string{"clientposts", pub.CNPJ}, ":"), pub.TimeStp, pub.Content)
	conn.Send("ZREMRANGEBYRANK", strings.Join([]string{"clientposts", pub.CNPJ}, ":"), "0", "0")
	//Althoug just update 2 field, still need to initialize the hset for the client manually with all the field
	//in a client sign in method later
	conn.Send("HMSET", strings.Join([]string{"client", pub.CNPJ}, ":"), "lastmessage", pub.Content, "timestamp", pub.TimeStp)
	_, err = conn.Do("EXEC")
	if err != nil {
		return nil, fmt.Errorf("Error in transaction: %v", err)
	}
	return &publish.Done{Status: 200}, nil
}

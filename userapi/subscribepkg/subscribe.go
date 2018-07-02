package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
	redis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"strings"
)

//Server - interface handler
type Server struct{}

//SubscribeRequest - grpc function implementation
func (s *Server) SubscribeRequest(ctx context.Context, subblock *subscribe.SubscribeBlock) (*subscribe.Done, error) {
	conn, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		return nil, fmt.Errorf("Error in db:%v", err)
	}
	defer conn.Close()
	arg1 := strings.Join([]string{"usersubs", subblock.OAuth}, ":")
	//Now the string in the key will hold name:cnpj
	arg2 := strings.Join([]string{subblock.Name, subblock.CNPJ}, ":")
	_, err = conn.Do("SADD", arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("Error in db: %v", err)
	}
	return &subscribe.Done{Status: 200}, nil
}

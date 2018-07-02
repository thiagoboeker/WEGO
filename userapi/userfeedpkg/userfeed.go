package userfeedpkg

import (
	"fmt"
	userfeed "github.com/WEGOAPP/userapi/protos/userfeed"
	redis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	"strings"
)

//Server - Interface Handler
type Server struct{}

//ExtractSubs - Extract the subs from the reply interface and send via the chan one by one
//this is excentialy for the concurrency restrictions for the redis conn so one by one they're fetch
func (s *Server) ExtractSubs(reply []interface{}, ch chan string, er chan error) {
	for _, v := range reply {
		switch v.(type) {
		case []interface{}:
			s, err := redis.ByteSlices(v, nil)
			if err != nil {
				er <- err
			} else {
				for _, str := range s {
					ch <- string(str)
				}
			}
		}
	}
	close(ch)
	close(er)
}

//ExtractFeed - Extract the feed from the client string, split in the ':' char to get the name in index 0
// and the cnpj in index 1
func (s *Server) ExtractFeed(client string, conn *redis.Conn) (*userfeed.Sub, error) {
	name := strings.Split(client, ":")
	cnpj := name[1]
	clientkey := strings.Join([]string{"client", cnpj}, ":")
	//variable to return
	var fblock userfeed.Sub
	fblock.Name = name[0]
	fblock.CNPJ = cnpj
	//reply from redis
	reply, err := redis.Values((*conn).Do("HMGET", clientkey, "lastmessage", "timestamp"))
	if err != nil {
		return nil, fmt.Errorf("Error fetching message")
	}
	//Pass the var from the reply to the returning struct
	_, err = redis.Scan(reply, &fblock.LastMessage, &fblock.LMtimestamp)
	if err != nil {
		return nil, fmt.Errorf("Error scaning for msgs")
	}
	return &fblock, nil
}

//UserSubs - RPC function to fetch the subs and the generate the main feed
func (s *Server) UserSubs(ctx context.Context, user *userfeed.User) (*userfeed.MainFeed, error) {
	conn, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		return nil, fmt.Errorf("Error in db: %v", err)
	}
	defer conn.Close()
	//Get the subs
	reply, err := redis.Values(conn.Do("SSCAN", strings.Join([]string{"usersubs", user.OAuth}, ":"), "0"))
	if err != nil {
		return nil, fmt.Errorf("Error in db: %v", err)
	}
	//two chanels to pass the client string and any error
	clientsch := make(chan string)
	erch := make(chan error)

	//Extracts the subs and sends via channel back here to main scope
	go s.ExtractSubs(reply, clientsch, erch)

	//variable to return
	var MainFeed userfeed.MainFeed

	//Range over the chanel and extract the feed from the client
	for client := range clientsch {
		fblock, err := s.ExtractFeed(client, &conn)
		if err != nil {
			fblock.LastMessage = "Error"
		}
		MainFeed.Feed = append(MainFeed.Feed, fblock)
	}
	return &MainFeed, nil

}

//GetFeedHistory - TODO
func (s *Server) GetFeedHistory(ctx context.Context, cnpj *userfeed.ClientCNPJ) (*userfeed.FeedHistory, error) {
	return &userfeed.FeedHistory{
		History: []*userfeed.FeedBlock{
			&userfeed.FeedBlock{
				Message:   "OK",
				TimeStamp: 12341231,
			},
		},
	}, nil
}

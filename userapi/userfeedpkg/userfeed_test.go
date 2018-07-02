package userfeedpkg

import (
	"fmt"
	userfeed "github.com/WEGOAPP/userapi/protos/userfeed"
	redis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

/*
func TestServer(t *testing.T) {
	t.Run("grpc", func(t *testing.T) {
		t.Parallel()
		mainctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		Run(mainctx, ":5000")
	})

	time.Sleep(3000 * time.Millisecond)

	t.Run("client", func(t *testing.T) {
		t.Parallel()
		conn, err := grpc.Dial(":5000", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		var client userfeed.UserFeedClient
		client = userfeed.NewUserFeedClient(conn)
		subgroup, err := client.GetUserFeed(context.Background(), &userfeed.User{
			OAuth: "1823c3u434",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(subgroup)
	})
}
*/

//Test for the UserSubs function
func TestSubFetch(t *testing.T) {
	t.Run("grpc", func(t *testing.T) {
		t.Parallel()
		mainctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		Run(mainctx, ":5000")
	})

	time.Sleep(3000)

	t.Run("client", func(t *testing.T) {
		t.Parallel()
		conn, err := grpc.Dial(":5000", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		var client userfeed.UserFeedClient
		client = userfeed.NewUserFeedClient(conn)
		reply, err := client.UserSubs(context.Background(), &userfeed.User{
			OAuth: "123abc",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	})
}

//Test for the ExtractFeed func
func TestMessageFetch(t *testing.T) {
	conn, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		log.Fatal(err)
	}
	var result struct {
		LastMessage string
		TimeStamp   int64
	}
	reply, err := redis.Values(conn.Do("HMGET", "client:54321", "lastmessage", "timestamp"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = redis.Scan(reply, &result.LastMessage, &result.TimeStamp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}

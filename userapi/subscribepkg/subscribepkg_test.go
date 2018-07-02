package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
	redis "github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	//Runs grpc server in parallel
	t.Run("grpc", func(t *testing.T) {
		t.Parallel()
		mainctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		Run(mainctx, ":4010")
	})

	time.Sleep(3000 * time.Millisecond)

	//Runs the client and does the rpc
	t.Run("client", func(t *testing.T) {
		t.Parallel()
		var client subscribe.SubscribeServiceClient
		conn, err := grpc.Dial(":4010", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client = subscribe.NewSubscribeServiceClient(conn)
		done, err := client.SubscribeRequest(context.Background(), &subscribe.SubscribeBlock{
			OAuth: "123abc",
			CNPJ:  "741852",
			Name:  "Franguinho da Ilha",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(done)
	})
}

//Test for the sub fetch mechanism
func TestSubFetch(t *testing.T) {
	conn, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		log.Fatal(err)
	}
	reply, err := redis.Values(conn.Do("SSCAN", "usersubs:123abc", "0"))
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan string)
	extract := func(reply []interface{}, ch chan string) {
		for _, v := range reply {
			switch v.(type) {
			case []interface{}:
				s, err := redis.ByteSlices(v, nil)
				if err != nil {
					log.Fatal(err)
				}
				for _, str := range s {
					ch <- string(str)
				}
			}
		}
		close(ch)
	}

	go extract(reply, ch)
	for sub := range ch {
		fmt.Println(sub)
	}
	fmt.Println("Done")

}

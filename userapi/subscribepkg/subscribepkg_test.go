package subscribepkg

import (
	"fmt"
	subscribe "github.com/WEGOAPP/userapi/protos/subscribe"
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
			CNPJ: "9192389128",
			Name: "Thiago",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(done)
	})
}

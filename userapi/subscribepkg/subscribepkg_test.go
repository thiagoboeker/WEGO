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
	t.Run("grpc", func(t *testing.T) {
		t.Parallel()
		if err := Run(":4010"); err != nil {
			log.Fatal(err)
		}
	})

	time.Sleep(3000 * time.Millisecond)

	t.Run("client", func(t *testing.T) {
		t.Parallel()
		var client subscribe.SubscribeServiceClient
		conn, err := grpc.Dial(":4010", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client = subscribe.NewSubscribeServiceClient(conn)
		done, err := client.SubscribeRequest(context.Background(), &subscribe.SubscribeBlock{
			CNPJ: "13206867703",
			Name: "Thiago",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(done)
	})
}

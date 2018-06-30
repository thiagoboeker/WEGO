package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	//Starts the grpc server to listen on port 4000
	t.Run("grpc", func(t *testing.T) {
		t.Parallel()
		mainctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		Run(mainctx, ":4000")
	})

	time.Sleep(3000 * time.Millisecond)

	//Starts the client
	t.Run("client", func(t *testing.T) {
		t.Parallel()
		fmt.Println("Starting client")
		var client publish.PublishServiceClient
		conn, err := grpc.Dial(":4000", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client = publish.NewPublishServiceClient(conn)

		done, err := client.PublishRequest(context.Background(), &publish.PublishBlock{
			CNPJ:    "193819883918s",
			Name:    "Thiago Boeker",
			Content: "Hello world",
			TimeStp: 132026568,
		})

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(done)
	})

}

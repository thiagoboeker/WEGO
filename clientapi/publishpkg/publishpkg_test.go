package publishpkg

import (
	"fmt"
	publish "github.com/WEGOAPP/clientapi/protos/publish"
	redis "github.com/garyburd/redigo/redis"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

//Test client publish mechanism
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
			CNPJ:    "741852",
			Name:    "Chili Beans",
			Content: "Ola pessoal! Promocao de oculos hoje",
			TimeStp: 132026900,
		})

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(done)
	})

}

//Test to ping the docker redis server
func TestRedisConn(t *testing.T) {
	_, err := redis.Dial("tcp", "192.168.99.100:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexao realizada")
}

package userfeedpkg

import (
	"fmt"
	userfeed "github.com/WEGOAPP/userapi/protos/userfeed"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

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

package userfeedpkg

import (
	"fmt"
	userfeed "github.com/WEGOAPP/userapi/protos/userfeed"
	"golang.org/x/net/context"
)

//Server - Interface Handler
type Server struct{}

//GetUserFeed - returns the user feed
func (s *Server) GetUserFeed(ctx context.Context, user *userfeed.User) (*userfeed.SubGroup, error) {
	fmt.Println(user)
	return &userfeed.SubGroup{
		Subs: []*userfeed.Sub{
			&userfeed.Sub{
				CNPJ:        "1283718273",
				Name:        "Thiago",
				LastMessage: "heelo people",
				LMtimestamp: 1217318,
			},
		},
	}, nil
}

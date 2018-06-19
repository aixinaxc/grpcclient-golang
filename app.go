package main

import (
	"google.golang.org/grpc"
	"fmt"
	"grpcclient/testg"
	"context"
	"io"
)

func main() {
	url := "39.104.167.28:8080"
	conn,err := grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		fmt.Println("error:",err)
	}
	client := testg.NewUserServerClient(conn)
	u := new(testg.User)
	u.UserId = "1"
	user,err := client.GetUserById(context.Background(),u)
	fmt.Println("user",user)

	empty := new(testg.Empty)
	us,err := client.GetListStream(context.Background(),empty)
	for {
		uuu,err := us.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(uuu)
		fmt.Println(uuu.UserName)
	}

	uss,err := client.SetUserStream(context.Background())
	uss.Send(u)
	uss.CloseAndRecv()
}


package jike_sdk_go

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var (
	client *Client
	ctx    = context.Background()
)

func TestMain(t *testing.M) {
	client = NewClient(WithDebug(true))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	loginReq := &LoginWithPhoneAndPasswordReq{
		AreaCode:          "+86",
		MobilePhoneNumber: phoneNumber,
		Password:          password,
	}
	_, err := client.LoginWithPhoneAndPassword(ctx, loginReq)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(t.Run())
}

func TestClient_LoginWithPhoneAndPassword(t *testing.T) {
	req := &LoginWithPhoneAndPasswordReq{
		AreaCode:          "+86",
		MobilePhoneNumber: phoneNumber,
		Password:          password,
	}
	rsp, err := client.LoginWithPhoneAndPassword(ctx, req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
	fmt.Printf("at: %s\n", client.xAccessToken)
	fmt.Printf("rt: %s\n", client.xRefreshToken)
}

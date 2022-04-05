package jike_sdk_go

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestClient_CreatePost(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	name := "pic.jpg"
	utReq, err := NewUploadTokenReq(name, PIC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	utRsp, err := client.UploadToken(ctx, utReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	uReq := &UploadReq{
		UploadToken: utRsp.UpToken,
		Filename:    name,
	}
	uRsp, err := client.Upload(ctx, uReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	cpReq := &CreatePostReq{
		PictureKeys:          []string{uRsp.Key},
		SyncToPersonalUpdate: true,
		SubmitToTopic:        "59e58bea89ee3f0016b4d2c6",
		Content:              "哈哈2",
	}
	cpRsp, err := client.CreatePost(context.Background(), cpReq)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", cpRsp)
}

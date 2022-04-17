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

func TestClient_GetPosts(t *testing.T) {
	rsp, err := client.GetPosts(context.Background(), &GetPostsReq{
		ID:      "6252e4f837f63b0010f1dd9b",
		UserRef: RecommendFeedPost,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetPosts2(t *testing.T) {
	recommendRsp, err := client.RecommendFeedList(context.Background(), &RecommendFeedListReq{Trigger: TriggerUser})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	feeds := recommendRsp.Data
	if len(feeds) == 0 {
		t.Logf("no feed")
		return
	}

	postsRsp, err := client.GetPosts(context.Background(), &GetPostsReq{
		ID:      feeds[0].Id,
		UserRef: RecommendFeedPost,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", postsRsp)
}

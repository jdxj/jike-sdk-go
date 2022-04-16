package jike_sdk_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_RecommendFeedList(t *testing.T) {
	rsp, err := client.RecommendFeedList(context.Background(), &RecommendFeedListReq{
		Trigger: TriggerUser,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("tm: %s\n", rsp.ToastMessage)
	fmt.Printf("lm: %+v\n", rsp.LoadMoreKey)

	for _, v := range rsp.Data {
		fmt.Printf("lives: %+v\n", v.Lives)
		fmt.Printf("type: %+v\n", v.Type)
		fmt.Printf("id: %+v\n", v.Id)
		fmt.Printf("rt: %+v\n", v.ReadTrackInfo)
		fmt.Printf("dm: %+v\n", v.DislikeMenu)
		fmt.Printf("user: %+v\n", v.User)
		fmt.Printf("topic: %+v\n", v.Topic)
		fmt.Printf("content: %+v\n", v.Content)
		fmt.Printf("lc: %+v\n", v.LikeCount)
		fmt.Printf("cc: %+v\n", v.CommentCount)
		fmt.Printf("ca: %+v\n", v.CreatedAt)
		fmt.Printf("pic: %+v\n", v.Pictures)
		fmt.Printf("poi: %+v\n", v.Poi)
		fmt.Printf("rc: %+v\n", v.RepostCount)
		fmt.Printf("li: %+v\n", v.LikeIcon)
		fmt.Printf("svt: %+v\n", v.SectionViewType)
		fmt.Printf("title: %+v\n", v.Title)
		fmt.Printf("itemc: %+v\n", v.ItemsCount)
		fmt.Printf("url: %+v\n", v.Url)
		fmt.Printf("uit: %+v\n", v.UrlsInText)
		fmt.Printf("sc: %+v\n", v.ShareCount)
		fmt.Println("----------------------------")
	}
}

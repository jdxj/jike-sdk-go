package jike_sdk_go

import (
	"context"
	"fmt"
	"time"
)

const (
	RecommendFeed = "/recommendFeed"
)

type TriggerType string

const (
	TriggerUser TriggerType = "user"
)

type FeedType string

const (
	LiveGroup         FeedType = "LIVE_GROUP"
	OriginalPost      FeedType = "ORIGINAL_POST"
	SectionHeader     FeedType = "SECTION_HEADER"
	SectionFooter     FeedType = "SECTION_FOOTER"
	RecommendUserFeed FeedType = "RECOMMEND_USER_FEED"
)

type RecommendFeedListReq struct {
	Trigger TriggerType `json:"trigger"`
}

type Live struct {
	Id        string    `json:"id"`
	User      User      `json:"user"`
	Status    string    `json:"status"`
	Topic     Topic     `json:"topic"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Picture   Picture   `json:"picture"`
}

type AdditionalInfo struct {
	StreamerId string `json:"streamerId"`
	TopicId    string `json:"topicId"`

	GrayGroup              string `json:"grayGroup"`
	InterestType           string `json:"interestType"`
	Keywords               string `json:"keywords"`
	PictureClassifications string `json:"pictureClassifications"`
	PosterId               string `json:"posterId"`
	RecallTag              string `json:"recallTag"`
	RecommendId            string `json:"recommendId"`
	RecommendTime          string `json:"recommendTime"`
}

type ReadTrackInfo struct {
	RecallPolicy   string         `json:"recallPolicy"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
	CurrentPage    string         `json:"currentPage"`
}

type LiveItem struct {
	Live          Live          `json:"live"`
	Reason        string        `json:"reason"`
	ReadTrackInfo ReadTrackInfo `json:"readTrackInfo"`
}

type Payload struct {
	Id            string        `json:"id"`
	FeedType      string        `json:"feedType"`
	Type          string        `json:"type"`
	Key           string        `json:"key"`
	Reason        string        `json:"reason"`
	ReadTrackInfo ReadTrackInfo `json:"readTrackInfo"`
}

type Reason struct {
	Key     string  `json:"key"`
	Text    string  `json:"text"`
	Payload Payload `json:"payload"`
}

type DislikeMenu struct {
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	Reasons  []Reason `json:"reasons"`
}

type Poi struct {
	PoiId            string    `json:"poiId"`
	Name             string    `json:"name"`
	Location         []float64 `json:"location"`
	Pname            string    `json:"pname"`
	Cityname         string    `json:"cityname"`
	FormattedAddress string    `json:"formattedAddress,omitempty"`
}

type UrlsInTextItem struct {
	Title       string `json:"title"`
	OriginalUrl string `json:"originalUrl"`
	Url         string `json:"url"`
	Type        string `json:"type"`
}

type Feed struct {
	Lives           []LiveItem       `json:"lives"`
	Type            FeedType         `json:"type"`
	Id              string           `json:"id,omitempty"`
	ReadTrackInfo   ReadTrackInfo    `json:"readTrackInfo"`
	DislikeMenu     DislikeMenu      `json:"dislikeMenu"`
	User            User             `json:"user"`
	Topic           Topic            `json:"topic"`
	Content         string           `json:"content,omitempty"`
	LikeCount       int              `json:"likeCount,omitempty"`
	CommentCount    int              `json:"commentCount,omitempty"`
	CreatedAt       time.Time        `json:"createdAt,omitempty"`
	Pictures        []Picture        `json:"pictures"`
	Poi             Poi              `json:"poi"`
	RepostCount     int              `json:"repostCount,omitempty"`
	LikeIcon        string           `json:"likeIcon,omitempty"`
	SectionViewType string           `json:"sectionViewType,omitempty"`
	Title           string           `json:"title,omitempty"`
	ItemsCount      int              `json:"itemsCount,omitempty"`
	Url             string           `json:"url,omitempty"`
	UrlsInText      []UrlsInTextItem `json:"urlsInText"`
	ShareCount      int              `json:"shareCount,omitempty"`
}

type LoadMoreKey struct {
	Page int `json:"page"`
}

type RecommendFeedListRsp struct {
	Data         []Feed      `json:"data"`
	ToastMessage string      `json:"toastMessage"`
	LoadMoreKey  LoadMoreKey `json:"loadMoreKey"`
}

func (c *Client) RecommendFeedList(ctx context.Context, req *RecommendFeedListReq) (*RecommendFeedListRsp, error) {
	rsp, err := c.apiR(ctx).
		SetBody(req).
		SetResult(&RecommendFeedListRsp{}).
		Post(apiBaseURL + apiVersion + RecommendFeed + "/list")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*RecommendFeedListRsp), nil
}

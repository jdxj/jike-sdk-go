package jike_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const (
	Posts = "/originalPosts"
)

type CreatePostReq struct {
	PictureKeys          []string `json:"pictureKeys"`
	SyncToPersonalUpdate bool     `json:"syncToPersonalUpdate"`
	SubmitToTopic        string   `json:"submitToTopic"`
	Content              string   `json:"content"`
}

type SquarePicture struct {
	Format       string `json:"format"`
	PicUrl       string `json:"picUrl"`
	MiddlePicUrl string `json:"middlePicUrl"`
	SmallPicUrl  string `json:"smallPicUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`

	LivePhoto json.RawMessage `json:"livePhoto"`
	Themes    Themes          `json:"themes"`
	Nft       json.RawMessage `json:"nft"`
}

type Tips struct {
	InDraft   string `json:"inDraft"`
	InComment string `json:"inComment"`
}

type Topic struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	Content          string `json:"content"`
	SubscribersCount int    `json:"subscribersCount"`

	SquarePicture SquarePicture `json:"squarePicture"`

	BriefIntro    string `json:"briefIntro"`
	TopicType     string `json:"topicType"`
	OperateStatus string `json:"operateStatus"`
	IsValid       bool   `json:"isValid"`
	IsVerified    bool   `json:"isVerified"`
	TopicId       int    `json:"topicId"`
	IsSearchable  bool   `json:"isSearchable"`
	LikeIcon      string `json:"likeIcon"`
	MessagePrefix string `json:"messagePrefix"`

	InternalTags json.RawMessage `json:"internalTags"`
	CustomTags   json.RawMessage `json:"customTags"`

	AuditStatus string `json:"auditStatus"`

	NewCategory   json.RawMessage `json:"newCategory"`
	InvolvedUsers json.RawMessage `json:"involvedUsers"`

	EntryTab string `json:"entryTab"`

	Tabs             json.RawMessage `json:"tabs"`
	RectanglePicture json.RawMessage `json:"rectanglePicture"`

	PictureUrl               string    `json:"pictureUrl"`
	ThumbnailUrl             string    `json:"thumbnailUrl"`
	SubscribedStatusRawValue int       `json:"subscribedStatusRawValue"`
	SubscribedAt             time.Time `json:"subscribedAt"`
	Ref                      string    `json:"ref"`
	IsDreamTopic             bool      `json:"isDreamTopic"`
	IsAnonymous              bool      `json:"isAnonymous"`
	EnablePictureComments    bool      `json:"enablePictureComments"`
	EnablePictureWatermark   bool      `json:"enablePictureWatermark"`
	TimeForRank              string    `json:"timeForRank"`
	LastMessagePostTime      time.Time `json:"lastMessagePostTime"`
	CreatedAt                string    `json:"createdAt"`
	UpdatedAt                string    `json:"updatedAt"`
	SubscribersTextSuffix    string    `json:"subscribersTextSuffix"`
	SubscribersName          string    `json:"subscribersName"`
	FriendsAlsoSubscribe     string    `json:"friendsAlsoSubscribe"`

	Maintainer json.RawMessage `json:"maintainer"`

	IsUserTopicAdmin bool `json:"isUserTopicAdmin"`

	ActivitySection  json.RawMessage `json:"activitySection"`
	ActivitySections json.RawMessage `json:"activitySections"`
	Tips             Tips            `json:"tips"`
	ToppingArea      json.RawMessage `json:"toppingArea"`

	InShortcuts   bool   `json:"inShortcuts"`
	PreferSection string `json:"preferSection"`

	RelatedHashtags json.RawMessage `json:"relatedHashtags"`

	Intro                       string    `json:"intro"`
	SquarePostUpdateTime        time.Time `json:"squarePostUpdateTime"`
	SubscribersAction           string    `json:"subscribersAction"`
	ApproximateSubscribersCount string    `json:"approximateSubscribersCount"`
	SubscribersDescription      string    `json:"subscribersDescription"`
	IsCommentForbidden          bool      `json:"isCommentForbidden"`
	BotCount                    int       `json:"botCount"`
	RecentPost                  string    `json:"recentPost"`
	Source                      string    `json:"source"`
	EnableForUserPost           bool      `json:"enableForUserPost"`
}

type Picture struct {
	Key          string  `json:"key"`
	ThumbnailUrl string  `json:"thumbnailUrl"`
	SmallPicUrl  string  `json:"smallPicUrl"`
	MiddlePicUrl string  `json:"middlePicUrl"`
	PicUrl       string  `json:"picUrl"`
	Format       string  `json:"format"`
	CropperPosX  float64 `json:"cropperPosX"`
	CropperPosY  float64 `json:"cropperPosY"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`

	WatermarkPicUrl string `json:"watermarkPicUrl,omitempty"`
}

type WmpShare struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type Rollouts struct {
	WmpShare WmpShare `json:"wmpShare"`
}

type ScrollingSubtitle struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type PostStatus string

const (
	Normal PostStatus = "NORMAL"
)

type Pinned struct {
	PersonalUpdate bool `json:"personalUpdate"`
}

type Post struct {
	Id      string   `json:"id"`
	Type    FeedType `json:"type"`
	Content string   `json:"content"`

	UrlsInText []UrlsInTextItem `json:"urlsInText"`

	Status             PostStatus `json:"status"`
	IsCommentForbidden bool       `json:"isCommentForbidden"`
	LikeCount          int        `json:"likeCount"`
	CommentCount       int        `json:"commentCount"`
	RepostCount        int        `json:"repostCount"`
	ShareCount         int        `json:"shareCount"`

	Topic    Topic     `json:"topic"`
	Pictures []Picture `json:"pictures"`

	Collected bool `json:"collected"`

	CollectTime json.RawMessage `json:"collectTime"`
	User        User            `json:"user"`

	CreatedAt             time.Time `json:"createdAt"`
	IsFeatured            bool      `json:"isFeatured"`
	EnablePictureComments bool      `json:"enablePictureComments"`
	Repostable            bool      `json:"repostable"`

	Rollouts           Rollouts            `json:"rollouts"`
	ScrollingSubtitles []ScrollingSubtitle `json:"scrollingSubtitles"`

	Poi                  Poi           `json:"poi"`
	IsSuppressed         bool          `json:"isSuppressed"`
	Pinned               Pinned        `json:"pinned"`
	ShouldShowCommentTip bool          `json:"shouldShowCommentTip"`
	ReadTrackInfo        ReadTrackInfo `json:"readTrackInfo"`
}

type CreatePostRsp struct {
	Success bool   `json:"success"`
	Toast   string `json:"toast"`
	Data    Post   `json:"data"`
}

func (c *Client) CreatePost(ctx context.Context, req *CreatePostReq) (*CreatePostRsp, error) {
	rsp, err := c.apiR(ctx).
		SetBody(req).
		SetResult(&CreatePostRsp{}).
		Post(apiBaseURL + apiVersion + Posts + "/create")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*CreatePostRsp), nil
}

type UserRef string

const (
	RecommendFeedPost UserRef = "RECOMMEND_FEED_POST"
)

type GetPostsReq struct {
	ID      string
	UserRef UserRef
}

type GetPostsRsp struct {
	Data Post `json:"data"`
}

func (c *Client) GetPosts(ctx context.Context, req *GetPostsReq) (*GetPostsRsp, error) {
	rsp, err := c.apiR(ctx).
		SetQueryParams(map[string]string{
			"id":      req.ID,
			"userRef": string(req.UserRef),
		}).
		SetResult(&GetPostsRsp{}).
		Get(apiBaseURL + apiVersion + Posts + "/get")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*GetPostsRsp), nil
}

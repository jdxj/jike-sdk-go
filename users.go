package jike_sdk_go

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const (
	Users = "/users"
)

type LoginWithPhoneAndPasswordReq struct {
	AreaCode          string `json:"areaCode"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
	Password          string `json:"password"`
}

type AvatarImage struct {
	ThumbnailUrl string `json:"thumbnailUrl"`
	SmallPicUrl  string `json:"smallPicUrl"`
	MiddlePicUrl string `json:"middlePicUrl"`
	PicUrl       string `json:"picUrl"`
	Format       string `json:"format"`
}

type StatsCount struct {
	TopicSubscribed            int `json:"topicSubscribed"`
	TopicCreated               int `json:"topicCreated"`
	FollowedCount              int `json:"followedCount"`
	FollowingCount             int `json:"followingCount"`
	HighlightedPersonalUpdates int `json:"highlightedPersonalUpdates"`
	Liked                      int `json:"liked"`
	RespectedCount             int `json:"respectedCount"`
}

type BackgroundImage struct {
	PicUrl string `json:"picUrl"`
}

type Preferences struct {
	PersonalizedRecommendationOn  bool            `json:"personalizedRecommendationOn"`
	AutoPlayVideo                 bool            `json:"autoPlayVideo"`
	TopicTagGuideOn               bool            `json:"topicTagGuideOn"`
	DailyNotificationOn           bool            `json:"dailyNotificationOn"`
	FollowedNotificationOn        bool            `json:"followedNotificationOn"`
	ReplyNotificationAllowGroup   string          `json:"replyNotificationAllowGroup"`
	LikeNotificationOn            bool            `json:"likeNotificationOn"`
	RespectNotificationOn         bool            `json:"respectNotificationOn"`
	MentionNotificationOn         bool            `json:"mentionNotificationOn"`
	LiveNotificationOn            bool            `json:"liveNotificationOn"`
	SubscribeWeatherForecast      bool            `json:"subscribeWeatherForecast"`
	PrivateTopicSubscribe         bool            `json:"privateTopicSubscribe"`
	UndiscoverableByPhoneNumber   bool            `json:"undiscoverableByPhoneNumber"`
	SaveDataUsageMode             bool            `json:"saveDataUsageMode"`
	TopicPushSettings             string          `json:"topicPushSettings"`
	DebugLogOn                    bool            `json:"debugLogOn"`
	Env                           string          `json:"env"`
	SyncCommentToPersonalActivity bool            `json:"syncCommentToPersonalActivity"`
	RepostWithComment             bool            `json:"repostWithComment"`
	EnablePrivateChat             bool            `json:"enablePrivateChat"`
	BlockStrangerPoke             bool            `json:"blockStrangerPoke"`
	EnablePictureWatermark        bool            `json:"enablePictureWatermark"`
	EnableOperationPush           bool            `json:"enableOperationPush"`
	EnableChatSound               bool            `json:"enableChatSound"`
	HideSubscribeTab              bool            `json:"hideSubscribeTab"`
	OpenMessageTabOnLaunch        bool            `json:"openMessageTabOnLaunch"`
	TabOnLaunch                   string          `json:"tabOnLaunch"`
	UndiscoverableByWeiboUser     json.RawMessage `json:"undiscoverableByWeiboUser"`
	PaidMarket                    json.RawMessage `json:"paidMarket"`
	FollowingUpdatesSortBy        string          `json:"followingUpdatesSortBy"`
}

type QqUserInfo struct {
	Id           string `json:"id"`
	ExternalName string `json:"externalName"`
}

type School struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Major          string `json:"major"`
	EnrollmentYear int    `json:"enrollmentYear"`
	PrivacyType    string `json:"privacyType"`
}

type RelationshipState struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ProfileTag struct {
	PicUrl string `json:"picUrl,omitempty"`
	Text   string `json:"text"`
	Type   string `json:"type"`
}

type RestrictedAvatarChange struct {
	NextChangeDate string `json:"nextChangeDate"`
	Limits         int    `json:"limits"`
}

type RestrictedNameChange struct {
	NextChangeDate string `json:"nextChangeDate"`
	Limits         int    `json:"limits"`
}

type Dark struct {
	PicUrl string `json:"picUrl"`
}

type Themes struct {
	Dark Dark `json:"dark"`
}

type Sponsor struct {
	PicUrl string `json:"picUrl"`
	Themes Themes `json:"themes"`
}

type Decorations struct {
	Sponsor Sponsor `json:"sponsor"`
}

type RefRemark struct {
	Type  string `json:"type"`
	RefId string `json:"refId"`
}

type ProfileVisitInfo struct {
	TodayCount    int   `json:"todayCount"`
	LatestVisitor *User `json:"latestVisitor"`
}

type TrailingIcon struct {
	PicUrl  string  `json:"picUrl"`
	Url     string  `json:"url"`
	Picture Picture `json:"picture"`
}

type Medal struct {
	PicUrl      string `json:"picUrl"`
	Url         string `json:"url,omitempty"`
	Name        string `json:"name"`
	GotMedalAt  string `json:"gotMedalAt"`
	BadgePicUrl string `json:"badgePicUrl,omitempty"`
}

type User struct {
	Id            string    `json:"id"`
	Username      string    `json:"username"`
	ScreenName    string    `json:"screenName"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsVerified    bool      `json:"isVerified"`
	VerifyMessage string    `json:"verifyMessage"`
	BriefIntro    string    `json:"briefIntro"`

	AvatarImage AvatarImage `json:"avatarImage"`
	Decorations Decorations `json:"decorations"`

	ProfileImageUrl string `json:"profileImageUrl"`

	StatsCount StatsCount `json:"statsCount"`

	IsBannedForever bool `json:"isBannedForever"`
	IsSponsor       bool `json:"isSponsor"`

	BackgroundImage BackgroundImage `json:"backgroundImage"`

	Bio    string `json:"bio"`
	Gender string `json:"gender"`

	Preferences Preferences `json:"preferences"`

	IsBetaUser        bool   `json:"isBetaUser"`
	UsernameSet       bool   `json:"usernameSet"`
	WechatOpenId      string `json:"wechatOpenId"`
	WechatUnionId     string `json:"wechatUnionId"`
	QqOpenId          string `json:"qqOpenId"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
	AreaCode          string `json:"areaCode"`
	Birthday          string `json:"birthday"`

	QqUserInfo QqUserInfo `json:"qqUserInfo"`

	TopicRoles json.RawMessage `json:"topicRoles"`

	School School `json:"school"`

	Zodiac string `json:"zodiac"`

	RelationshipState RelationshipState `json:"relationshipState"`

	ShowRespect bool    `json:"showRespect"`
	Medals      []Medal `json:"medals"`

	ProfileTags []ProfileTag `json:"profileTags"`

	BackgroundMaskColor string `json:"backgroundMaskColor"`
	IsLoginUser         bool   `json:"isLoginUser"`
	IsBanned            bool   `json:"isBanned"`
	UserHasPosted       bool   `json:"userHasPosted"`
	HasStories          bool   `json:"hasStories"`

	RestrictedAvatarChange RestrictedAvatarChange `json:"restrictedAvatarChange"`
	RestrictedNameChange   RestrictedNameChange   `json:"restrictedNameChange"`

	IsDefaultScreenName bool `json:"isDefaultScreenName"`

	ProfileVisitInfo ProfileVisitInfo `json:"profileVisitInfo"`

	City             string         `json:"city"`
	Country          string         `json:"country"`
	Province         string         `json:"province"`
	Ref              string         `json:"ref"`
	SponsorExpiresAt time.Time      `json:"sponsorExpiresAt"`
	RefRemark        RefRemark      `json:"refRemark"`
	TrailingIcons    []TrailingIcon `json:"trailingIcons"`
	Following        bool           `json:"following"`
	Blocking         bool           `json:"blocking"`
	Muting           bool           `json:"muting"`
	Watching         bool           `json:"watching"`
	IsWatching       bool           `json:"isWatching"`
	StoryStatus      string         `json:"storyStatus"`
	Respected        bool           `json:"respected"`
}

type LoginWithPhoneAndPasswordRsp struct {
	IsRegister      bool            `json:"isRegister"`
	User            User            `json:"user"`
	EnabledFeatures []string        `json:"enabledFeatures"`
	AgreedProtocol  json.RawMessage `json:"agreedProtocol"`
}

// todo: 何时刷新 token?
func (c *Client) LoginWithPhoneAndPassword(ctx context.Context, req *LoginWithPhoneAndPasswordReq) (
	*LoginWithPhoneAndPasswordRsp, error) {
	rsp, err := c.apiR(ctx).
		SetBody(req).
		SetResult(&LoginWithPhoneAndPasswordRsp{}).
		Post(apiBaseURL + apiVersion + Users + "/loginWithPhoneAndPassword")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	loginRsp := rsp.Result().(*LoginWithPhoneAndPasswordRsp)
	c.xAccessToken = rsp.Header().Get("x-jike-access-token")
	c.xRefreshToken = rsp.Header().Get("x-jike-refresh-token")
	c.xRequestID = rsp.Header().Get("x-request-id")
	return loginRsp, nil
}

type GetProfileReq struct {
	Username string
}

type RelationUsers struct {
	Message string `json:"message"`
	Users   []User `json:"users"`
}

type GetProfileRsp struct {
	User            User          `json:"user"`
	RelationMessage string        `json:"relationMessage"`
	RelationUsers   RelationUsers `json:"relationUsers"`
	Invisible       bool          `json:"invisible"`
}

func (c *Client) GetProfile(ctx context.Context, req *GetProfileReq) (*GetProfileRsp, error) {
	rsp, err := c.apiR(ctx).
		SetQueryParam("username", req.Username).
		SetResult(&GetProfileRsp{}).
		Get(apiBaseURL + apiVersion + Users + "/profile")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*GetProfileRsp), nil
}

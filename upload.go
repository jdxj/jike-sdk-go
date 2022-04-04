package jike_sdk_go

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type UploadType = string

const (
	PIC UploadType = "PIC"
)

func NewUploadTokenReq(name string, ut UploadType) (*UploadTokenReq, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return nil, err
	}

	req := &UploadTokenReq{
		md5:        fmt.Sprintf("%x", hash.Sum(nil)),
		uploadType: ut,
	}
	return req, nil
}

type UploadTokenReq struct {
	md5        string
	uploadType UploadType
}

func (utr *UploadTokenReq) Map() map[string]string {
	return map[string]string{
		"md5":        utr.md5,
		"uploadType": utr.uploadType,
	}
}

type UploadTokenRsp struct {
	UpToken string `json:"uptoken"`
}

func (c *Client) UploadToken(ctx context.Context, req *UploadTokenReq) (*UploadTokenRsp, error) {
	rsp, err := c.uploadR(ctx).
		SetQueryParams(req.Map()).
		SetResult(&UploadTokenRsp{}).
		Get(uploadBaseURL + uploadVersion + "/upload/token")
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*UploadTokenRsp), nil
}

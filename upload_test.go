package jike_sdk_go

import (
	"context"
	"fmt"
	"testing"
)

func TestNewUploadTokenReq(t *testing.T) {
	req, err := NewUploadTokenReq("README.md", PIC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", req)
}

func TestClient_UploadToken(t *testing.T) {
	req, err := NewUploadTokenReq(picPath, PIC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	rsp, err := client.UploadToken(context.Background(), req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

package jike_sdk_go

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"path/filepath"
	"testing"
	"time"
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

func TestMockUploadUserAgent(t *testing.T) {
	res := mockUploadUserAgent(ut)
	fmt.Printf("res: %s\n", res)
}

func TestMockUploadFilename(t *testing.T) {
	ext := filepath.Ext("abc.txt")
	res := mockUploadFilename(ext)
	fmt.Printf("res: %s\n", res)
}

func TestClient_Upload(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	name := "pic.jpg"
	utr, err := NewUploadTokenReq(name, PIC)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	utp, err := client.UploadToken(ctx, utr)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	rsp, err := client.Upload(ctx, &UploadReq{
		UploadToken: utp.UpToken,
		Filename:    name,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestRandRead(t *testing.T) {
	buf := make([]byte, 4)
	for i := 0; i < 10; i++ {
		_, _ = rand.Read(buf)
		res := binary.LittleEndian.Uint32(buf)
		fmt.Printf("%d\n", res)
	}
}

package jike_sdk_go

import (
	"fmt"
	"strings"
	"time"
)

var (
	mockAPIHost    = "api.ruguoapp.com"
	mockAPIHeaders = map[string]string{
		"os":                       "Android",
		"os-version":               "30",
		"app-version":              "7.31.1",
		"app-buildno":              "2475",
		"manufacturer":             "Xiaomi",
		"model":                    "Redmi K20 Pro",
		"resolution":               "1080x2208",
		"market":                   "update",
		"applicationid":            "com.ruguoapp.jike",
		"x-jike-app-id":            "XeITUMa6kGKF",
		"x-jike-device-id":         "bd9f2d90-dd53-47d7-85af-42366c77c9af",
		"king-card-status":         "unknown",
		"app-permissions":          "4",
		"x-jike-device-properties": "{\"uuid\":\"bd9f2d90-dd53-47d7-85af-42366c77c9af\",\"android_id\":\"e7cd7c21655a7d6b\",\"oaid\":\"\",\"vaid\":\"\",\"aaid\":\"\"}",
		"accept-encoding":          "gzip",
		"user-agent":               "okhttp/4.9.2",
	}
	mockUploadTokenHost = "upload.ruguoapp.com"

	mockUploadHeaders = map[string]string{
		"Host":            "upload.qiniup.com",
		"accept-encoding": "gzip",
	}
)

func mockUploadUserAgent(ut string) string {
	os := "QiniuAndroid/8.4.2"
	osVer := "11"
	manufacturer := "Xiaomi/Redmi"
	model := "K20 Pro/30"
	ts := time.Now().UnixMicro()
	utHead := ut[:16]
	return fmt.Sprintf("%s (%s; %s %s; %d; %s)",
		os, osVer, manufacturer, model, ts, utHead)
}

func mockUploadFilename(ext string) string {
	prefix := "clean"
	suffix := GenRandString(9)
	ts := time.Now().Format("2006-01-02-15-04-05.000")
	ts = strings.ReplaceAll(ts, ".", "-")
	return fmt.Sprintf("%s-%s-%s%s", prefix, ts, suffix, ext)
}

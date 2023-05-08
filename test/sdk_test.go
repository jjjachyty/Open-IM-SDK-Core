package test

import (
	"encoding/json"
	"fmt"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"testing"
	"time"
)

func Test_xxx(t *testing.T) {
	var cf sdk_struct.IMConfig
	cf.ApiAddr = APIADDR
	cf.Platform = 1
	cf.WsAddr = WSADDR
	cf.DataDir = "./"
	cf.LogLevel = LogLevel
	cf.ObjectStorage = "minio"
	cf.IsCompression = true
	cf.IsExternalExtensions = true

	var s string
	b, _ := json.Marshal(cf)
	s = string(b)
	fmt.Println(s)
	var testinit testInitLister

	operationID := utils.OperationIDGenerator()
	if !open_im_sdk.InitSDK(&testinit, operationID, s) {
		log.Error("", "InitSDK failed")
		return
	}
	x := &XBase{}
	open_im_sdk.Login(x, operationID, me, "")

	DotestGetGroupMemberList()
	time.Sleep(time.Second * 2)
}

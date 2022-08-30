package tests

import (
	"os"
	"testing"

	tbkApi "github.com/bububa/opentaobao/api/tbk"
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model"
	tbkModel "github.com/bububa/opentaobao/model/tbk"
)

func TestJSONFormatHmac(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	secret := os.Getenv("SECRET")
	testURL := os.Getenv("TEST_TBK_LINK")
	clt := core.NewSDKClient(appKey, secret)
	clt.SetDebug(true)
	clt.SetAPIFormat(model.JSON)
	clt.SetSignMethod(model.HMAC)
	req := tbkModel.NewTaobaoTbkTpwdCreateRequest()
	req.SetUrl(testURL)
	req.SetText("test tpwd create")
	resp, err := tbkApi.TaobaoTbkTpwdCreate(clt, req, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("RESP: %+v\n", *resp.Data)
}

func TestJSONFormatMd5(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	secret := os.Getenv("SECRET")
	testURL := os.Getenv("TEST_TBK_LINK")
	clt := core.NewSDKClient(appKey, secret)
	clt.SetDebug(true)
	clt.SetAPIFormat(model.JSON)
	clt.SetSignMethod(model.MD5)
	req := tbkModel.NewTaobaoTbkTpwdCreateRequest()
	req.SetUrl(testURL)
	req.SetText("test tpwd create")
	resp, err := tbkApi.TaobaoTbkTpwdCreate(clt, req, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("RESP: %+v\n", *resp.Data)
}

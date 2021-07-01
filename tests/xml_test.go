package tests

import (
	"os"
	"testing"

	tbkApi "github.com/bububa/opentaobao/api/tbk"
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model"
	tbkModel "github.com/bububa/opentaobao/model/tbk"
)

func TestXMLFormatHmac(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	secret := os.Getenv("SECRET")
	clt := core.NewSDKClient(appKey, secret)
	clt.SetDebug(true)
	clt.SetAPIFormat(model.XML)
	clt.SetSignMethod(model.HMAC)
	req := tbkModel.NewTaobaoTbkItemInfoGetRequest()
	req.SetNumIids("617361236670")
	req.SetPlatform(2)
	resp, err := tbkApi.TaobaoTbkItemInfoGet(clt, req, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("RESP: %+v\n", *resp)
}

func TestXMLFormatMd5(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	secret := os.Getenv("SECRET")
	clt := core.NewSDKClient(appKey, secret)
	clt.SetDebug(true)
	clt.SetAPIFormat(model.XML)
	clt.SetSignMethod(model.MD5)
	req := tbkModel.NewTaobaoTbkItemInfoGetRequest()
	req.SetNumIids("617361236670")
	req.SetPlatform(2)
	resp, err := tbkApi.TaobaoTbkItemInfoGet(clt, req, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("RESP: %+v\n", *resp)
}

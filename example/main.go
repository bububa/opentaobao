package main

import (
	"log"
	"os"

	tbkApi "github.com/bububa/opentaobao/api/tbk"
	"github.com/bububa/opentaobao/core"
	tbkModel "github.com/bububa/opentaobao/model/tbk"
)

func main() {
	appKey := os.Getenv("APPKEY")
	secret := os.Getenv("SECRET")
	clt := core.NewSDKClient(appKey, secret)
	clt.SetDebug(true)
	clt.SetAPIFormat("xml")
	clt.SetSignMethod("hmac")
	req := tbkModel.NewTaobaoTbkItemInfoGetRequest()
	req.SetNumIids("617361236670")
	req.SetPlatform(2)
	resp, err := tbkApi.TaobaoTbkItemInfoGet(clt, req, "")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("RESP: %+v\n", *resp)
}

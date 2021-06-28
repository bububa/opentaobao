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
	req := tbkModel.NewTaobaoTbkTpwdCreateRequest()
	req.SetUrl("https://s.click.taobao.com/GKQEZlu")
	req.SetText("test tpwd create")
	resp, err := tbkApi.TaobaoTbkTpwdCreate(clt, req, "")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("RESP: %+v\n", *resp.Data)
}

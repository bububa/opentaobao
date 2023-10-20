package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmssignnamedelete 淘宝短信签名删除
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
func Taobaojstsmssignnamedelete(clt *core.SDKClient, req *jst.TaobaojstsmssignnamedeleteAPIRequest, session string) (*jst.TaobaojstsmssignnamedeleteAPIResponse, error) {
	var resp jst.TaobaojstsmssignnamedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

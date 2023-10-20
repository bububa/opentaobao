package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmssignnamequery 淘宝短信签名查询
// taobao.jst.sms.signname.query
//
// 淘宝短信签名查询
func Taobaojstsmssignnamequery(clt *core.SDKClient, req *jst.TaobaojstsmssignnamequeryAPIRequest, session string) (*jst.TaobaojstsmssignnamequeryAPIResponse, error) {
	var resp jst.TaobaojstsmssignnamequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmssignnamecreate 淘宝短信签名创建接口
// taobao.jst.sms.signname.create
//
// 聚石塔短信签名创建接口
func Taobaojstsmssignnamecreate(clt *core.SDKClient, req *jst.TaobaojstsmssignnamecreateAPIRequest, session string) (*jst.TaobaojstsmssignnamecreateAPIResponse, error) {
	var resp jst.TaobaojstsmssignnamecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

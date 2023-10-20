package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmsoaidmessagesend 基于OAID的短信发送接口
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
func Taobaojstsmsoaidmessagesend(clt *core.SDKClient, req *jst.TaobaojstsmsoaidmessagesendAPIRequest, session string) (*jst.TaobaojstsmsoaidmessagesendAPIResponse, error) {
	var resp jst.TaobaojstsmsoaidmessagesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

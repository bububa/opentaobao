package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstminiappopenidmessagesend 单个openId用户短信发送
// taobao.jst.miniapp.openid.message.send
//
// 单个openId用户短信发送
func Taobaojstminiappopenidmessagesend(clt *core.SDKClient, req *jst.TaobaojstminiappopenidmessagesendAPIRequest, session string) (*jst.TaobaojstminiappopenidmessagesendAPIResponse, error) {
	var resp jst.TaobaojstminiappopenidmessagesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

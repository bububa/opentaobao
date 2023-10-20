package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstMiniappOpenidMessageSend 单个openId用户短信发送
// taobao.jst.miniapp.openid.message.send
//
// 单个openId用户短信发送
func TaobaoJstMiniappOpenidMessageSend(clt *core.SDKClient, req *jst.TaobaoJstMiniappOpenidMessageSendAPIRequest, resp *jst.TaobaoJstMiniappOpenidMessageSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

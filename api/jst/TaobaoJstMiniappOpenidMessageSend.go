package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstMiniappOpenidMessageSend 单个openId用户短信发送
// taobao.jst.miniapp.openid.message.send
//
// 单个openId用户短信发送
func TaobaoJstMiniappOpenidMessageSend(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstMiniappOpenidMessageSendAPIRequest, resp *jst.TaobaoJstMiniappOpenidMessageSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

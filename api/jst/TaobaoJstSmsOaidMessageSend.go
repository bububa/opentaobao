package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsOaidMessageSend 基于OAID的短信发送接口
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
func TaobaoJstSmsOaidMessageSend(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstSmsOaidMessageSendAPIRequest, resp *jst.TaobaoJstSmsOaidMessageSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

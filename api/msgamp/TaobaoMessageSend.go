package msgamp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/msgamp"
)

// TaobaoMessageSend 消息发送
// taobao.message.send
//
// 消息发送接口
func TaobaoMessageSend(ctx context.Context, clt *core.SDKClient, req *msgamp.TaobaoMessageSendAPIRequest, resp *msgamp.TaobaoMessageSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

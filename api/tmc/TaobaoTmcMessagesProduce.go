package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcMessagesProduce 批量发送消息
// taobao.tmc.messages.produce
//
// 批量发送消息
func TaobaoTmcMessagesProduce(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcMessagesProduceAPIRequest, resp *tmc.TaobaoTmcMessagesProduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

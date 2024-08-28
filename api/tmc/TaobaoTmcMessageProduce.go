package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcMessageProduce 发布单条消息
// taobao.tmc.message.produce
//
// 发布单条消息
func TaobaoTmcMessageProduce(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcMessageProduceAPIRequest, resp *tmc.TaobaoTmcMessageProduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

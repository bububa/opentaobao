package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoQimenEventProduce 发出奇门事件
// taobao.qimen.event.produce
//
// 当订单被处理时，用于通知奇门系统。
func TaobaoQimenEventProduce(ctx context.Context, clt *core.SDKClient, req *util.TaobaoQimenEventProduceAPIRequest, resp *util.TaobaoQimenEventProduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

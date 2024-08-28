package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleOrderDummySend 闲鱼无需物流发货
// alibaba.idle.order.dummy.send
//
// 适用于电子卡券等虚拟商品不需要物流的商品发货。
func AlibabaIdleOrderDummySend(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleOrderDummySendAPIRequest, resp *idle.AlibabaIdleOrderDummySendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

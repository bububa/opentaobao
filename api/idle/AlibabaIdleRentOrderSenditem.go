package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderSenditem 确认发货
// alibaba.idle.rent.order.senditem
//
// 确认发货
func AlibabaIdleRentOrderSenditem(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentOrderSenditemAPIRequest, resp *idle.AlibabaIdleRentOrderSenditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

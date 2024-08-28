package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderReceiveitem 确认签收
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
func AlibabaIdleRentOrderReceiveitem(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentOrderReceiveitemAPIRequest, resp *idle.AlibabaIdleRentOrderReceiveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

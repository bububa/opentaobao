package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderQuery 查询订单
// alibaba.idle.rent.order.query
//
// 查询订单信息
func AlibabaIdleRentOrderQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentOrderQueryAPIRequest, resp *idle.AlibabaIdleRentOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

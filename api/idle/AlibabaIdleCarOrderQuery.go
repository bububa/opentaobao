package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleCarOrderQuery 二手车寄卖查询订单接口
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
func AlibabaIdleCarOrderQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleCarOrderQueryAPIRequest, resp *idle.AlibabaIdleCarOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

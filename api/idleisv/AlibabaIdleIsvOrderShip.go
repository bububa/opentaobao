package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvOrderShip 闲鱼订单服务商物流发货
// alibaba.idle.isv.order.ship
//
// 闲鱼开放平台服务商订单发货接口
func AlibabaIdleIsvOrderShip(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderShipAPIRequest, resp *idleisv.AlibabaIdleIsvOrderShipAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

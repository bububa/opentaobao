package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvOrderAdjustprice 闲鱼服务商订单价格修改接口
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
func AlibabaIdleIsvOrderAdjustprice(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderAdjustpriceAPIRequest, resp *idleisv.AlibabaIdleIsvOrderAdjustpriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

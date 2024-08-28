package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvOrderClose 服务商闲鱼卖家主动关闭订单
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
func AlibabaIdleIsvOrderClose(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvOrderCloseAPIRequest, resp *idleisv.AlibabaIdleIsvOrderCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

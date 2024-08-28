package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleConsignmentiiOrderPerform 寄卖V2订单履约
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
func AlibabaIdleConsignmentiiOrderPerform(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleConsignmentiiOrderPerformAPIRequest, resp *idle.AlibabaIdleConsignmentiiOrderPerformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

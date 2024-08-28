package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleConsignmentOrderPerform 帮卖订单履约
// alibaba.idle.consignment.order.perform
//
// 帮卖订单履约，回收商同步订单信息，驱动交易流转
func AlibabaIdleConsignmentOrderPerform(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleConsignmentOrderPerformAPIRequest, resp *idle.AlibabaIdleConsignmentOrderPerformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

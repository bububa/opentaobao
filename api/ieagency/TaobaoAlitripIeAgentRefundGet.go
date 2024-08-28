package ieagency

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundGet 获取退票申请详情
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
func TaobaoAlitripIeAgentRefundGet(ctx context.Context, clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundGetAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

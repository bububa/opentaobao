package ieagency

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundRefundmoney 确认退款
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
func TaobaoAlitripIeAgentRefundRefundmoney(ctx context.Context, clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

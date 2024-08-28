package ieagency

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundNewReceive 商家退票受理申请(对外)
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
func TaobaoAlitripIeAgentRefundNewReceive(ctx context.Context, clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundNewReceiveAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundNewReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

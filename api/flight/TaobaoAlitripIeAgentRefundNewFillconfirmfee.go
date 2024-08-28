package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripIeAgentRefundNewFillconfirmfee 新模型-回填申请单费用
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
func TaobaoAlitripIeAgentRefundNewFillconfirmfee(ctx context.Context, clt *core.SDKClient, req *flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest, resp *flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

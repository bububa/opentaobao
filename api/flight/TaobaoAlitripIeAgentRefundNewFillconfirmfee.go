package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripIeAgentRefundNewFillconfirmfee 新模型-回填申请单费用
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
func TaobaoAlitripIeAgentRefundNewFillconfirmfee(clt *core.SDKClient, req *flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest, resp *flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

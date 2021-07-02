package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundAgree 同意退票
// taobao.alitrip.ie.agent.refund.agree
//
// 卖家同意买家退票申请
func TaobaoAlitripIeAgentRefundAgree(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundAgreeAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundAgreeAPIResponse, error) {
	var resp ieagency.TaobaoAlitripIeAgentRefundAgreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundGet 获取退票申请详情
// taobao.alitrip.ie.agent.refund.get
//
// 获取退票申请详情
func TaobaoAlitripIeAgentRefundGet(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundGetAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundGetAPIResponse, error) {
	var resp ieagency.TaobaoAlitripIeAgentRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

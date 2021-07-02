package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundNewMultiplerefunds 补退接口
// taobao.alitrip.ie.agent.refund.new.multiplerefunds
//
// 1. 补退接口， 可以进行多次退款
func TaobaoAlitripIeAgentRefundNewMultiplerefunds(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponse, error) {
	var resp ieagency.TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

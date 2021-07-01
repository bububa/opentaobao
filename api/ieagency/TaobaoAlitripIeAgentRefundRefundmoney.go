package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

/* TaobaoAlitripIeAgentRefundRefundmoney
确认退款
taobao.alitrip.ie.agent.refund.refundmoney

卖家进行退款操作 */
func TaobaoAlitripIeAgentRefundRefundmoney(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse, error) {
	var resp ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

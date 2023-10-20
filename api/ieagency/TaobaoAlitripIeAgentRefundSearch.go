package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundSearch 卖家查询退票申请
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
func TaobaoAlitripIeAgentRefundSearch(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundSearchAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

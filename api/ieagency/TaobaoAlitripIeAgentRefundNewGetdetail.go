package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundNewGetdetail 查询申请单详情(新版)
// taobao.alitrip.ie.agent.refund.new.getdetail
//
// 查询申请单详情
func TaobaoAlitripIeAgentRefundNewGetdetail(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

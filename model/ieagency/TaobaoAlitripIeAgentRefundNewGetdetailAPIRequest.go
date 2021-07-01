package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest
查询申请单详情(新版) API请求
taobao.alitrip.ie.agent.refund.new.getdetail

查询申请单详情 */
type TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq
}

// New

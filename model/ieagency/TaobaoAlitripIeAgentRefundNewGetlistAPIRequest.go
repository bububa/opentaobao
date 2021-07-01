package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundNewGetlistAPIRequest
新查询退票申请单列表 API请求
taobao.alitrip.ie.agent.refund.new.getlist

查询商家国际机票退票申请单列表 */
type TaobaoAlitripIeAgentRefundNewGetlistAPIRequest struct {
	model.Params
	// 列表查询
	_paramRefundOrderQueryListRq *RefundOrderQueryListRq
}

// New

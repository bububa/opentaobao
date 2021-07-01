package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest
新模型-回填申请单费用 API请求
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用 */
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq
}

// New

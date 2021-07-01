package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellRefundRefuseAPIRequest
销售退票拒绝 API请求
alitrip.agent.flight.sell.refund.refuse

销售退票拒绝 */
type AlitripAgentFlightSellRefundRefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国内国际标识
	_domesticIntl int64
	// 拒绝原因
	_refuseReason string
}

// New

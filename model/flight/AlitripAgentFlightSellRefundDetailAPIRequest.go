package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellRefundDetailAPIRequest
销售退票单详情 API请求
alitrip.agent.flight.sell.refund.detail

销售退票单详情 */
type AlitripAgentFlightSellRefundDetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国际国内标识
	_domesticIntl int64
}

// New

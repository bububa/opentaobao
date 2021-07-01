package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellRefundApproveAPIRequest
销售退票确认 API请求
alitrip.agent.flight.sell.refund.approve

销售退票确认 */
type AlitripAgentFlightSellRefundApproveAPIRequest struct {
	model.Params
	// 入参
	_param *RefundApproveRequestDto
}

// New

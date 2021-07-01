package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellRefundListAPIRequest
销售退票单列表 API请求
alitrip.agent.flight.sell.refund.list

销售退票单列表 */
type AlitripAgentFlightSellRefundListAPIRequest struct {
	model.Params
	// 请求对象
	_param *RefundListRequestDto
}

// New

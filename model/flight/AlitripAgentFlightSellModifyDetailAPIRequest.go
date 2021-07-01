package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyDetailAPIRequest
销售改签详情 API请求
alitrip.agent.flight.sell.modify.detail

销售改签详情 */
type AlitripAgentFlightSellModifyDetailAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国内国际标识
	_domesticIntl int64
}

// New

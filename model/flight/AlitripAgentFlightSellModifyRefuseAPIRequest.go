package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyRefuseAPIRequest
销售改签拒绝 API请求
alitrip.agent.flight.sell.modify.refuse

销售改签拒绝 */
type AlitripAgentFlightSellModifyRefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 国际国内标识
	_domesticIntl int64
	// 拒绝原因
	_refuseReason string
}

// New

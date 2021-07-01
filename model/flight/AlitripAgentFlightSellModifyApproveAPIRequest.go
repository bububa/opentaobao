package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyApproveAPIRequest
销售改签确认 API请求
alitrip.agent.flight.sell.modify.approve

销售改签确认 */
type AlitripAgentFlightSellModifyApproveAPIRequest struct {
	model.Params
	// 入参对象
	_param *ModifyApproveRequestDto
}

// New

package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyListAPIRequest
销售改签单列表 API请求
alitrip.agent.flight.sell.modify.list

销售改签单列表 */
type AlitripAgentFlightSellModifyListAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyListRequestDto
}

// New

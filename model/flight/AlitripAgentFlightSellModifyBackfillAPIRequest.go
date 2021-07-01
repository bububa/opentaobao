package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyBackfillAPIRequest
销售改签回填 API请求
alitrip.agent.flight.sell.modify.backfill

销售改签回填 */
type AlitripAgentFlightSellModifyBackfillAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyBackFillRequestDto
}

// New

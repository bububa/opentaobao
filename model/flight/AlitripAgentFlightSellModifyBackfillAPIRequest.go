package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellmodifybackfillAPIRequest 销售改签回填 API请求
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
type AlitripagentflightsellmodifybackfillAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyBackFillRequestDto
}

// NewAlitripagentflightsellmodifybackfillRequest 初始化AlitripagentflightsellmodifybackfillAPIRequest对象
func NewAlitripagentflightsellmodifybackfillRequest() *AlitripagentflightsellmodifybackfillAPIRequest {
	return &AlitripagentflightsellmodifybackfillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellmodifybackfillAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.backfill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellmodifybackfillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellmodifybackfillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlitripagentflightsellmodifybackfillAPIRequest) SetParam(_param *ModifyBackFillRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellmodifybackfillAPIRequest) GetParam() *ModifyBackFillRequestDto {
	return r._param
}

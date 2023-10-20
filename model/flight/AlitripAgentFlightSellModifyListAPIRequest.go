package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellmodifylistAPIRequest 销售改签单列表 API请求
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
type AlitripagentflightsellmodifylistAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyListRequestDto
}

// NewAlitripagentflightsellmodifylistRequest 初始化AlitripagentflightsellmodifylistAPIRequest对象
func NewAlitripagentflightsellmodifylistRequest() *AlitripagentflightsellmodifylistAPIRequest {
	return &AlitripagentflightsellmodifylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellmodifylistAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellmodifylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellmodifylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlitripagentflightsellmodifylistAPIRequest) SetParam(_param *ModifyListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellmodifylistAPIRequest) GetParam() *ModifyListRequestDto {
	return r._param
}

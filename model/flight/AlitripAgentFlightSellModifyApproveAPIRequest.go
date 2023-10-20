package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellmodifyapproveAPIRequest 销售改签确认 API请求
// alitrip.agent.flight.sell.modify.approve
//
// 销售改签确认
type AlitripagentflightsellmodifyapproveAPIRequest struct {
	model.Params
	// 入参对象
	_param *ModifyApproveRequestDto
}

// NewAlitripagentflightsellmodifyapproveRequest 初始化AlitripagentflightsellmodifyapproveAPIRequest对象
func NewAlitripagentflightsellmodifyapproveRequest() *AlitripagentflightsellmodifyapproveAPIRequest {
	return &AlitripagentflightsellmodifyapproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellmodifyapproveAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellmodifyapproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellmodifyapproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripagentflightsellmodifyapproveAPIRequest) SetParam(_param *ModifyApproveRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellmodifyapproveAPIRequest) GetParam() *ModifyApproveRequestDto {
	return r._param
}

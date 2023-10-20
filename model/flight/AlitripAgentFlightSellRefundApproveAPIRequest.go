package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundapproveAPIRequest 销售退票确认 API请求
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
type AlitripagentflightsellrefundapproveAPIRequest struct {
	model.Params
	// 入参
	_param *RefundApproveRequestDto
}

// NewAlitripagentflightsellrefundapproveRequest 初始化AlitripagentflightsellrefundapproveAPIRequest对象
func NewAlitripagentflightsellrefundapproveRequest() *AlitripagentflightsellrefundapproveAPIRequest {
	return &AlitripagentflightsellrefundapproveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellrefundapproveAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellrefundapproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellrefundapproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlitripagentflightsellrefundapproveAPIRequest) SetParam(_param *RefundApproveRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellrefundapproveAPIRequest) GetParam() *RefundApproveRequestDto {
	return r._param
}

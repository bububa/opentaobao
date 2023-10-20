package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundlistAPIRequest 销售退票单列表 API请求
// alitrip.agent.flight.sell.refund.list
//
// 销售退票单列表
type AlitripagentflightsellrefundlistAPIRequest struct {
	model.Params
	// 请求对象
	_param *RefundListRequestDto
}

// NewAlitripagentflightsellrefundlistRequest 初始化AlitripagentflightsellrefundlistAPIRequest对象
func NewAlitripagentflightsellrefundlistRequest() *AlitripagentflightsellrefundlistAPIRequest {
	return &AlitripagentflightsellrefundlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellrefundlistAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellrefundlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellrefundlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求对象
func (r *AlitripagentflightsellrefundlistAPIRequest) SetParam(_param *RefundListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellrefundlistAPIRequest) GetParam() *RefundListRequestDto {
	return r._param
}

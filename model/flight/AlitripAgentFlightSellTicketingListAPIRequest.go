package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellticketinglistAPIRequest 销售出票列表 API请求
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
type AlitripagentflightsellticketinglistAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingListRequestDto
}

// NewAlitripagentflightsellticketinglistRequest 初始化AlitripagentflightsellticketinglistAPIRequest对象
func NewAlitripagentflightsellticketinglistRequest() *AlitripagentflightsellticketinglistAPIRequest {
	return &AlitripagentflightsellticketinglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellticketinglistAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellticketinglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellticketinglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripagentflightsellticketinglistAPIRequest) SetParam(_param *TicketingListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellticketinglistAPIRequest) GetParam() *TicketingListRequestDto {
	return r._param
}

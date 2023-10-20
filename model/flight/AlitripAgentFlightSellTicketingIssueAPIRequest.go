package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellticketingissueAPIRequest 销售出票 API请求
// alitrip.agent.flight.sell.ticketing.issue
//
// 销售出票
type AlitripagentflightsellticketingissueAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingIssueRequestDto
}

// NewAlitripagentflightsellticketingissueRequest 初始化AlitripagentflightsellticketingissueAPIRequest对象
func NewAlitripagentflightsellticketingissueRequest() *AlitripagentflightsellticketingissueAPIRequest {
	return &AlitripagentflightsellticketingissueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightsellticketingissueAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightsellticketingissueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightsellticketingissueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripagentflightsellticketingissueAPIRequest) SetParam(_param *TicketingIssueRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripagentflightsellticketingissueAPIRequest) GetParam() *TicketingIssueRequestDto {
	return r._param
}

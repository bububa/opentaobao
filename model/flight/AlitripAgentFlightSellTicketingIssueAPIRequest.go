package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingIssueAPIRequest 销售出票 API请求
// alitrip.agent.flight.sell.ticketing.issue
//
// 销售出票
type AlitripAgentFlightSellTicketingIssueAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingIssueRequestDto
}

// NewAlitripAgentFlightSellTicketingIssueRequest 初始化AlitripAgentFlightSellTicketingIssueAPIRequest对象
func NewAlitripAgentFlightSellTicketingIssueRequest() *AlitripAgentFlightSellTicketingIssueAPIRequest {
	return &AlitripAgentFlightSellTicketingIssueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripAgentFlightSellTicketingIssueAPIRequest) SetParam(_param *TicketingIssueRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetParam() *TicketingIssueRequestDto {
	return r._param
}

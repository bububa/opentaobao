package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票 API请求
alitrip.agent.flight.sell.ticketing.issue

销售出票
*/
type AlitripAgentFlightSellTicketingIssueRequest struct {
    model.Params
    // 入参对象
    _param   *TicketingIssueRequestDto
}

// 初始化AlitripAgentFlightSellTicketingIssueRequest对象
func NewAlitripAgentFlightSellTicketingIssueRequest() *AlitripAgentFlightSellTicketingIssueRequest{
    return &AlitripAgentFlightSellTicketingIssueRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingIssueRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.issue"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlitripAgentFlightSellTicketingIssueRequest) SetParam(_param *TicketingIssueRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellTicketingIssueRequest) GetParam() *TicketingIssueRequestDto {
    return r._param
}

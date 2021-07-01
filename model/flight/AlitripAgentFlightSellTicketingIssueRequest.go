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
type AlitripAgentFlightSellTicketingIssueAPIRequest struct {
    model.Params
    // 入参对象
    _param   *TicketingIssueRequestDTO
}

// 初始化AlitripAgentFlightSellTicketingIssueAPIRequest对象
func NewAlitripAgentFlightSellTicketingIssueRequest() *AlitripAgentFlightSellTicketingIssueAPIRequest{
    return &AlitripAgentFlightSellTicketingIssueAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.issue"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlitripAgentFlightSellTicketingIssueAPIRequest) SetParam(_param *TicketingIssueRequestDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellTicketingIssueAPIRequest) GetParam() *TicketingIssueRequestDTO {
    return r._param
}

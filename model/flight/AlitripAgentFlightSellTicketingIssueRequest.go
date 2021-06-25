package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票 APIRequest
alitrip.agent.flight.sell.ticketing.issue

销售出票
*/
type AlitripAgentFlightSellTicketingIssueRequest struct {
    model.Params

    // 入参对象
    param   *TicketingIssueRequestDto 

}

func NewAlitripAgentFlightSellTicketingIssueRequest() *AlitripAgentFlightSellTicketingIssueRequest{
    return &AlitripAgentFlightSellTicketingIssueRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellTicketingIssueRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.issue"
}

func (r AlitripAgentFlightSellTicketingIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellTicketingIssueRequest) SetParam(param *TicketingIssueRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellTicketingIssueRequest) GetParam() *TicketingIssueRequestDto {
    return r.param
}


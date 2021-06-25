package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票列表 APIRequest
alitrip.agent.flight.sell.ticketing.list

销售出票列表
*/
type AlitripAgentFlightSellTicketingListRequest struct {
    model.Params

    // 入参对象
    param   *TicketingListRequestDto 

}

func NewAlitripAgentFlightSellTicketingListRequest() *AlitripAgentFlightSellTicketingListRequest{
    return &AlitripAgentFlightSellTicketingListRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellTicketingListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.list"
}

func (r AlitripAgentFlightSellTicketingListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellTicketingListRequest) SetParam(param *TicketingListRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellTicketingListRequest) GetParam() *TicketingListRequestDto {
    return r.param
}


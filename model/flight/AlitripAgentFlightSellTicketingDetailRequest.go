package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票详情 APIRequest
alitrip.agent.flight.sell.ticketing.detail

销售出票详情
*/
type AlitripAgentFlightSellTicketingDetailRequest struct {
    model.Params

    // 国内国际标识
    domesticIntl   int64 

    // 飞猪订单号
    orderId   string 

}

func NewAlitripAgentFlightSellTicketingDetailRequest() *AlitripAgentFlightSellTicketingDetailRequest{
    return &AlitripAgentFlightSellTicketingDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellTicketingDetailRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.ticketing.detail"
}

func (r AlitripAgentFlightSellTicketingDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellTicketingDetailRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

func (r AlitripAgentFlightSellTicketingDetailRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}

func (r *AlitripAgentFlightSellTicketingDetailRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripAgentFlightSellTicketingDetailRequest) GetOrderId() string {
    return r.orderId
}


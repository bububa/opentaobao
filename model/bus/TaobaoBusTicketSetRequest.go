package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出票接口 APIRequest
taobao.bus.ticket.set

提供给汽车票商家出票使用
*/
type TaobaoBusTicketSetRequest struct {
    model.Params

    // 系统自动生成
    ticketParams   *B2BBookOrderRq 

}

func NewTaobaoBusTicketSetRequest() *TaobaoBusTicketSetRequest{
    return &TaobaoBusTicketSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTicketSetRequest) GetApiMethodName() string {
    return "taobao.bus.ticket.set"
}

func (r TaobaoBusTicketSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTicketSetRequest) SetTicketParams(ticketParams *B2BBookOrderRq) error {
    r.ticketParams = ticketParams
    r.Set("ticket_params", ticketParams)
    return nil
}

func (r TaobaoBusTicketSetRequest) GetTicketParams() *B2BBookOrderRq {
    return r.ticketParams
}


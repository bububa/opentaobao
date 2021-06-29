package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口invalidTicket APIRequest
alibaba.damai.mev.open.invalidticket

开放接口 使票无效
*/
type AlibabaDamaiMevOpenInvalidticketRequest struct {
    model.Params

    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam 

}

func NewAlibabaDamaiMevOpenInvalidticketRequest() *AlibabaDamaiMevOpenInvalidticketRequest{
    return &AlibabaDamaiMevOpenInvalidticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenInvalidticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.invalidticket"
}

func (r AlibabaDamaiMevOpenInvalidticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenInvalidticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

func (r AlibabaDamaiMevOpenInvalidticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}


package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口changeTicket APIRequest
alibaba.damai.mev.open.changeticket

开放接口 换票
*/
type AlibabaDamaiMevOpenChangeticketRequest struct {
    model.Params

    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam 

}

func NewAlibabaDamaiMevOpenChangeticketRequest() *AlibabaDamaiMevOpenChangeticketRequest{
    return &AlibabaDamaiMevOpenChangeticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenChangeticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.changeticket"
}

func (r AlibabaDamaiMevOpenChangeticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenChangeticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

func (r AlibabaDamaiMevOpenChangeticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}


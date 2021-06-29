package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口resetTicket APIRequest
alibaba.damai.mev.open.resetticket

开放接口重打票
*/
type AlibabaDamaiMevOpenResetticketRequest struct {
    model.Params

    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam 

}

func NewAlibabaDamaiMevOpenResetticketRequest() *AlibabaDamaiMevOpenResetticketRequest{
    return &AlibabaDamaiMevOpenResetticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenResetticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.resetticket"
}

func (r AlibabaDamaiMevOpenResetticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenResetticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

func (r AlibabaDamaiMevOpenResetticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}


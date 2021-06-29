package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口unlockTicket APIRequest
alibaba.damai.mev.open.unlockticket

开放接口 解锁票单
*/
type AlibabaDamaiMevOpenUnlockticketRequest struct {
    model.Params

    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam 

}

func NewAlibabaDamaiMevOpenUnlockticketRequest() *AlibabaDamaiMevOpenUnlockticketRequest{
    return &AlibabaDamaiMevOpenUnlockticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenUnlockticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.unlockticket"
}

func (r AlibabaDamaiMevOpenUnlockticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenUnlockticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

func (r AlibabaDamaiMevOpenUnlockticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}


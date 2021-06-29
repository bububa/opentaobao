package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口lockTicket APIRequest
alibaba.damai.mev.open.lockticket

开放接口 冻结票单
*/
type AlibabaDamaiMevOpenLockticketRequest struct {
    model.Params

    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam 

}

func NewAlibabaDamaiMevOpenLockticketRequest() *AlibabaDamaiMevOpenLockticketRequest{
    return &AlibabaDamaiMevOpenLockticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenLockticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.lockticket"
}

func (r AlibabaDamaiMevOpenLockticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenLockticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

func (r AlibabaDamaiMevOpenLockticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}


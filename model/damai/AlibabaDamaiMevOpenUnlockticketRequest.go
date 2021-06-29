package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口unlockTicket API请求
alibaba.damai.mev.open.unlockticket

开放接口 解锁票单
*/
type AlibabaDamaiMevOpenUnlockticketRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenUnlockticketRequest对象
func NewAlibabaDamaiMevOpenUnlockticketRequest() *AlibabaDamaiMevOpenUnlockticketRequest{
    return &AlibabaDamaiMevOpenUnlockticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenUnlockticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.unlockticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenUnlockticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenUnlockticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenUnlockticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}

package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口withdrawTicket API请求
alibaba.damai.mev.open.withdrawticket

开放接口退票
*/
type AlibabaDamaiMevOpenWithdrawticketRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenWithdrawticketRequest对象
func NewAlibabaDamaiMevOpenWithdrawticketRequest() *AlibabaDamaiMevOpenWithdrawticketRequest{
    return &AlibabaDamaiMevOpenWithdrawticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenWithdrawticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.withdrawticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenWithdrawticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenWithdrawticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenWithdrawticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}

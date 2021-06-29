package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口changeTicket API请求
alibaba.damai.mev.open.changeticket

开放接口 换票
*/
type AlibabaDamaiMevOpenChangeticketRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenChangeticketRequest对象
func NewAlibabaDamaiMevOpenChangeticketRequest() *AlibabaDamaiMevOpenChangeticketRequest{
    return &AlibabaDamaiMevOpenChangeticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenChangeticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.changeticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenChangeticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenChangeticketRequest) SetTicketIdOpenParam(ticketIdOpenParam *TicketIdOpenParam) error {
    r.ticketIdOpenParam = ticketIdOpenParam
    r.Set("ticket_id_open_param", ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenChangeticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r.ticketIdOpenParam
}

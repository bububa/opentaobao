package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口invalidTicket API请求
alibaba.damai.mev.open.invalidticket

开放接口 使票无效
*/
type AlibabaDamaiMevOpenInvalidticketRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    _ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenInvalidticketRequest对象
func NewAlibabaDamaiMevOpenInvalidticketRequest() *AlibabaDamaiMevOpenInvalidticketRequest{
    return &AlibabaDamaiMevOpenInvalidticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenInvalidticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.invalidticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenInvalidticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenInvalidticketRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
    r._ticketIdOpenParam = _ticketIdOpenParam
    r.Set("ticket_id_open_param", _ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenInvalidticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r._ticketIdOpenParam
}

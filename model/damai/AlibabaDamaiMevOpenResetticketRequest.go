package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口resetTicket API请求
alibaba.damai.mev.open.resetticket

开放接口重打票
*/
type AlibabaDamaiMevOpenResetticketRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    _ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenResetticketRequest对象
func NewAlibabaDamaiMevOpenResetticketRequest() *AlibabaDamaiMevOpenResetticketRequest{
    return &AlibabaDamaiMevOpenResetticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenResetticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.resetticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenResetticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenResetticketRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
    r._ticketIdOpenParam = _ticketIdOpenParam
    r.Set("ticket_id_open_param", _ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenResetticketRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r._ticketIdOpenParam
}

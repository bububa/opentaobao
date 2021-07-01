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
type AlibabaDamaiMevOpenResetticketAPIRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    _ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenResetticketAPIRequest对象
func NewAlibabaDamaiMevOpenResetticketRequest() *AlibabaDamaiMevOpenResetticketAPIRequest{
    return &AlibabaDamaiMevOpenResetticketAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.resetticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenResetticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
    r._ticketIdOpenParam = _ticketIdOpenParam
    r.Set("ticket_id_open_param", _ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenResetticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r._ticketIdOpenParam
}

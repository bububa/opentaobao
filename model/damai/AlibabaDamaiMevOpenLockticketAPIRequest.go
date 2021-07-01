package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口lockTicket API请求
alibaba.damai.mev.open.lockticket

开放接口 冻结票单
*/
type AlibabaDamaiMevOpenLockticketAPIRequest struct {
    model.Params
    // 入参ticketIdOpenParam
    _ticketIdOpenParam   *TicketIdOpenParam
}

// 初始化AlibabaDamaiMevOpenLockticketAPIRequest对象
func NewAlibabaDamaiMevOpenLockticketRequest() *AlibabaDamaiMevOpenLockticketAPIRequest{
    return &AlibabaDamaiMevOpenLockticketAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.lockticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabaDamaiMevOpenLockticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
    r._ticketIdOpenParam = _ticketIdOpenParam
    r.Set("ticket_id_open_param", _ticketIdOpenParam)
    return nil
}

// TicketIdOpenParam Getter
func (r AlibabaDamaiMevOpenLockticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
    return r._ticketIdOpenParam
}

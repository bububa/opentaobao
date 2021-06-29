package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出票接口 API请求
taobao.bus.ticket.set

提供给汽车票商家出票使用
*/
type TaobaoBusTicketSetRequest struct {
    model.Params
    // 系统自动生成
    _ticketParams   *B2BBookOrderRq
}

// 初始化TaobaoBusTicketSetRequest对象
func NewTaobaoBusTicketSetRequest() *TaobaoBusTicketSetRequest{
    return &TaobaoBusTicketSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTicketSetRequest) GetApiMethodName() string {
    return "taobao.bus.ticket.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTicketSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketParams Setter
// 系统自动生成
func (r *TaobaoBusTicketSetRequest) SetTicketParams(_ticketParams *B2BBookOrderRq) error {
    r._ticketParams = _ticketParams
    r.Set("ticket_params", _ticketParams)
    return nil
}

// TicketParams Getter
func (r TaobaoBusTicketSetRequest) GetTicketParams() *B2BBookOrderRq {
    return r._ticketParams
}

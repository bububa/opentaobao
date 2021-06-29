package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部确认 API请求
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
*/
type TaobaoLifeReservationItemOrderConfirmRequest struct {
    model.Params
    // 淘宝主单号
    _tradeNo   string
    // 凭证ID
    _ticketId   string
    // 审核类型，PASS-通过；REJECT-驳回
    _optType   string
}

// 初始化TaobaoLifeReservationItemOrderConfirmRequest对象
func NewTaobaoLifeReservationItemOrderConfirmRequest() *TaobaoLifeReservationItemOrderConfirmRequest{
    return &TaobaoLifeReservationItemOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationItemOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.life.reservation.item.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationItemOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r TaobaoLifeReservationItemOrderConfirmRequest) GetTradeNo() string {
    return r._tradeNo
}
// TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetTicketId(_ticketId string) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r TaobaoLifeReservationItemOrderConfirmRequest) GetTicketId() string {
    return r._ticketId
}
// OptType Setter
// 审核类型，PASS-通过；REJECT-驳回
func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetOptType(_optType string) error {
    r._optType = _optType
    r.Set("opt_type", _optType)
    return nil
}

// OptType Getter
func (r TaobaoLifeReservationItemOrderConfirmRequest) GetOptType() string {
    return r._optType
}

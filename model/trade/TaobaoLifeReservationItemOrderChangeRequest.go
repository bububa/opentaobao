package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部发起变更 API请求
taobao.life.reservation.item.order.change

生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
*/
type TaobaoLifeReservationItemOrderChangeRequest struct {
    model.Params
    // 淘宝主单号
    _tradeNo   string
    // 凭证ID
    _ticketId   string
    // 改期：MODIFY   取消：CANCEL
    _action   string
    // 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
    _reserveStartTime   string
    // 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
    _reserveEndTime   string
}

// 初始化TaobaoLifeReservationItemOrderChangeRequest对象
func NewTaobaoLifeReservationItemOrderChangeRequest() *TaobaoLifeReservationItemOrderChangeRequest{
    return &TaobaoLifeReservationItemOrderChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationItemOrderChangeRequest) GetApiMethodName() string {
    return "taobao.life.reservation.item.order.change"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationItemOrderChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetTradeNo() string {
    return r._tradeNo
}
// TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetTicketId(_ticketId string) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetTicketId() string {
    return r._ticketId
}
// Action Setter
// 改期：MODIFY   取消：CANCEL
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetAction() string {
    return r._action
}
// ReserveStartTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetReserveStartTime(_reserveStartTime string) error {
    r._reserveStartTime = _reserveStartTime
    r.Set("reserve_start_time", _reserveStartTime)
    return nil
}

// ReserveStartTime Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetReserveStartTime() string {
    return r._reserveStartTime
}
// ReserveEndTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetReserveEndTime(_reserveEndTime string) error {
    r._reserveEndTime = _reserveEndTime
    r.Set("reserve_end_time", _reserveEndTime)
    return nil
}

// ReserveEndTime Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetReserveEndTime() string {
    return r._reserveEndTime
}

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
    tradeNo   string
    // 凭证ID
    ticketId   string
    // 改期：MODIFY   取消：CANCEL
    action   string
    // 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
    reserveStartTime   string
    // 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
    reserveEndTime   string
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
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

// TradeNo Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetTradeNo() string {
    return r.tradeNo
}
// TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

// TicketId Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetTicketId() string {
    return r.ticketId
}
// Action Setter
// 改期：MODIFY   取消：CANCEL
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetAction() string {
    return r.action
}
// ReserveStartTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetReserveStartTime(reserveStartTime string) error {
    r.reserveStartTime = reserveStartTime
    r.Set("reserve_start_time", reserveStartTime)
    return nil
}

// ReserveStartTime Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetReserveStartTime() string {
    return r.reserveStartTime
}
// ReserveEndTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
func (r *TaobaoLifeReservationItemOrderChangeRequest) SetReserveEndTime(reserveEndTime string) error {
    r.reserveEndTime = reserveEndTime
    r.Set("reserve_end_time", reserveEndTime)
    return nil
}

// ReserveEndTime Getter
func (r TaobaoLifeReservationItemOrderChangeRequest) GetReserveEndTime() string {
    return r.reserveEndTime
}

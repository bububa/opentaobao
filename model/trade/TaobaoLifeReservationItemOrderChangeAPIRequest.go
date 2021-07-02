package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLifeReservationItemOrderChangeAPIRequest 生服购后预约单外部发起变更 API请求
// taobao.life.reservation.item.order.change
//
// 生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
type TaobaoLifeReservationItemOrderChangeAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
	// 改期：MODIFY   取消：CANCEL
	_action string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
	_reserveStartTime string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
	_reserveEndTime string
}

// NewTaobaoLifeReservationItemOrderChangeRequest 初始化TaobaoLifeReservationItemOrderChangeAPIRequest对象
func NewTaobaoLifeReservationItemOrderChangeRequest() *TaobaoLifeReservationItemOrderChangeAPIRequest {
	return &TaobaoLifeReservationItemOrderChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.item.order.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeNo is TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationItemOrderChangeAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetTicketId is TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationItemOrderChangeAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetAction is Action Setter
// 改期：MODIFY   取消：CANCEL
func (r *TaobaoLifeReservationItemOrderChangeAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetAction() string {
	return r._action
}

// SetReserveStartTime is ReserveStartTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
func (r *TaobaoLifeReservationItemOrderChangeAPIRequest) SetReserveStartTime(_reserveStartTime string) error {
	r._reserveStartTime = _reserveStartTime
	r.Set("reserve_start_time", _reserveStartTime)
	return nil
}

// GetReserveStartTime ReserveStartTime Getter
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetReserveStartTime() string {
	return r._reserveStartTime
}

// SetReserveEndTime is ReserveEndTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
func (r *TaobaoLifeReservationItemOrderChangeAPIRequest) SetReserveEndTime(_reserveEndTime string) error {
	r._reserveEndTime = _reserveEndTime
	r.Set("reserve_end_time", _reserveEndTime)
	return nil
}

// GetReserveEndTime ReserveEndTime Getter
func (r TaobaoLifeReservationItemOrderChangeAPIRequest) GetReserveEndTime() string {
	return r._reserveEndTime
}

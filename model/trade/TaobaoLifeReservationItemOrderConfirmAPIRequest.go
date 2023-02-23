package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLifeReservationItemOrderConfirmAPIRequest 生服购后预约单外部确认 API请求
// taobao.life.reservation.item.order.confirm
//
// 生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
type TaobaoLifeReservationItemOrderConfirmAPIRequest struct {
	model.Params
	// 凭证ID。与预约单号二选一，优先级低
	_ticketId string
	// 审核结果。PASS：通过；REJECT：驳回
	_optType string
	// 预约单号。与凭证ID二选一，优先级高
	_reservationOrderId string
}

// NewTaobaoLifeReservationItemOrderConfirmRequest 初始化TaobaoLifeReservationItemOrderConfirmAPIRequest对象
func NewTaobaoLifeReservationItemOrderConfirmRequest() *TaobaoLifeReservationItemOrderConfirmAPIRequest {
	return &TaobaoLifeReservationItemOrderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.item.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketId is TicketId Setter
// 凭证ID。与预约单号二选一，优先级低
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetOptType is OptType Setter
// 审核结果。PASS：通过；REJECT：驳回
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetOptType(_optType string) error {
	r._optType = _optType
	r.Set("opt_type", _optType)
	return nil
}

// GetOptType OptType Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetOptType() string {
	return r._optType
}

// SetReservationOrderId is ReservationOrderId Setter
// 预约单号。与凭证ID二选一，优先级高
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetReservationOrderId(_reservationOrderId string) error {
	r._reservationOrderId = _reservationOrderId
	r.Set("reservation_order_id", _reservationOrderId)
	return nil
}

// GetReservationOrderId ReservationOrderId Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetReservationOrderId() string {
	return r._reservationOrderId
}

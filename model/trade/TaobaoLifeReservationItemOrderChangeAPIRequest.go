package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolifereservationitemorderchangeAPIRequest 生服购后预约单外部发起变更 API请求
// taobao.life.reservation.item.order.change
//
// 生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
type TaobaolifereservationitemorderchangeAPIRequest struct {
	model.Params
	// 凭证ID。与预约单号二选一，优先级低
	_ticketId string
	// 改期：MODIFY ；取消：CANCEL；推进履约：PUSH_FULFILLMENT
	_action string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
	_reserveStartTime string
	// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
	_reserveEndTime string
	// 目标履约状态。 枚举值：PREPARING,PREPARED,PROCESSING,PROCESSED,DELIVERING,FINISHED。详以接入文档中描述的场景对接
	_targetFulfillmentStatus string
	// 预约单号。与凭证ID二选一，优先级高
	_reservationOrderId string
	// 扩展信息
	_extInfo *CommonKeyValue
}

// NewTaobaolifereservationitemorderchangeRequest 初始化TaobaolifereservationitemorderchangeAPIRequest对象
func NewTaobaolifereservationitemorderchangeRequest() *TaobaolifereservationitemorderchangeAPIRequest {
	return &TaobaolifereservationitemorderchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolifereservationitemorderchangeAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.item.order.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolifereservationitemorderchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolifereservationitemorderchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketId is TicketId Setter
// 凭证ID。与预约单号二选一，优先级低
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetAction is Action Setter
// 改期：MODIFY ；取消：CANCEL；推进履约：PUSH_FULFILLMENT
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetAction() string {
	return r._action
}

// SetReserveStartTime is ReserveStartTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定00:00
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetReserveStartTime(_reserveStartTime string) error {
	r._reserveStartTime = _reserveStartTime
	r.Set("reserve_start_time", _reserveStartTime)
	return nil
}

// GetReserveStartTime ReserveStartTime Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetReserveStartTime() string {
	return r._reserveStartTime
}

// SetReserveEndTime is ReserveEndTime Setter
// 改期必填，格式：yyyy-MM-dd HH:mm。时分固定23:59
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetReserveEndTime(_reserveEndTime string) error {
	r._reserveEndTime = _reserveEndTime
	r.Set("reserve_end_time", _reserveEndTime)
	return nil
}

// GetReserveEndTime ReserveEndTime Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetReserveEndTime() string {
	return r._reserveEndTime
}

// SetTargetFulfillmentStatus is TargetFulfillmentStatus Setter
// 目标履约状态。 枚举值：PREPARING,PREPARED,PROCESSING,PROCESSED,DELIVERING,FINISHED。详以接入文档中描述的场景对接
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetTargetFulfillmentStatus(_targetFulfillmentStatus string) error {
	r._targetFulfillmentStatus = _targetFulfillmentStatus
	r.Set("target_fulfillment_status", _targetFulfillmentStatus)
	return nil
}

// GetTargetFulfillmentStatus TargetFulfillmentStatus Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetTargetFulfillmentStatus() string {
	return r._targetFulfillmentStatus
}

// SetReservationOrderId is ReservationOrderId Setter
// 预约单号。与凭证ID二选一，优先级高
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetReservationOrderId(_reservationOrderId string) error {
	r._reservationOrderId = _reservationOrderId
	r.Set("reservation_order_id", _reservationOrderId)
	return nil
}

// GetReservationOrderId ReservationOrderId Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetReservationOrderId() string {
	return r._reservationOrderId
}

// SetExtInfo is ExtInfo Setter
// 扩展信息
func (r *TaobaolifereservationitemorderchangeAPIRequest) SetExtInfo(_extInfo *CommonKeyValue) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r TaobaolifereservationitemorderchangeAPIRequest) GetExtInfo() *CommonKeyValue {
	return r._extInfo
}

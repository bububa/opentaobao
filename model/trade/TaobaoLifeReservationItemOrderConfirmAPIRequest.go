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
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
	// 审核类型，PASS-通过；REJECT-驳回
	_optType string
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
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// Get TradeNo Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// Set is TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// Get TicketId Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetTicketId() string {
	return r._ticketId
}

// Set is OptType Setter
// 审核类型，PASS-通过；REJECT-驳回
func (r *TaobaoLifeReservationItemOrderConfirmAPIRequest) SetOptType(_optType string) error {
	r._optType = _optType
	r.Set("opt_type", _optType)
	return nil
}

// Get OptType Getter
func (r TaobaoLifeReservationItemOrderConfirmAPIRequest) GetOptType() string {
	return r._optType
}

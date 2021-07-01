package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLifeReservationTradeConsumeNoticeAPIRequest
生服购后预约外部核销 API请求
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销 */
type TaobaoLifeReservationTradeConsumeNoticeAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
}

// NewTaobaoLifeReservationTradeConsumeNoticeRequest 初始化TaobaoLifeReservationTradeConsumeNoticeAPIRequest对象
func NewTaobaoLifeReservationTradeConsumeNoticeRequest() *TaobaoLifeReservationTradeConsumeNoticeAPIRequest {
	return &TaobaoLifeReservationTradeConsumeNoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.trade.consume.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// Get TradeNo Getter
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// Set is TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// Get TicketId Getter
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetTicketId() string {
	return r._ticketId
}

package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolifereservationtradeconsumenoticeAPIRequest 生服购后预约外部核销 API请求
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
type TaobaolifereservationtradeconsumenoticeAPIRequest struct {
	model.Params
	// 淘宝主单号
	_tradeNo string
	// 凭证ID
	_ticketId string
}

// NewTaobaolifereservationtradeconsumenoticeRequest 初始化TaobaolifereservationtradeconsumenoticeAPIRequest对象
func NewTaobaolifereservationtradeconsumenoticeRequest() *TaobaolifereservationtradeconsumenoticeAPIRequest {
	return &TaobaolifereservationtradeconsumenoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolifereservationtradeconsumenoticeAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.trade.consume.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolifereservationtradeconsumenoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolifereservationtradeconsumenoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 淘宝主单号
func (r *TaobaolifereservationtradeconsumenoticeAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r TaobaolifereservationtradeconsumenoticeAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetTicketId is TicketId Setter
// 凭证ID
func (r *TaobaolifereservationtradeconsumenoticeAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r TaobaolifereservationtradeconsumenoticeAPIRequest) GetTicketId() string {
	return r._ticketId
}

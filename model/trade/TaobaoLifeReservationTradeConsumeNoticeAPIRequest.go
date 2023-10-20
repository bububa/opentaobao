package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLifeReservationTradeConsumeNoticeAPIRequest 生服购后预约外部核销 API请求
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) Reset() {
	r._tradeNo = ""
	r._ticketId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetApiMethodName() string {
	return "taobao.life.reservation.trade.consume.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetTicketId is TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r TaobaoLifeReservationTradeConsumeNoticeAPIRequest) GetTicketId() string {
	return r._ticketId
}

var poolTaobaoLifeReservationTradeConsumeNoticeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLifeReservationTradeConsumeNoticeRequest()
	},
}

// GetTaobaoLifeReservationTradeConsumeNoticeRequest 从 sync.Pool 获取 TaobaoLifeReservationTradeConsumeNoticeAPIRequest
func GetTaobaoLifeReservationTradeConsumeNoticeAPIRequest() *TaobaoLifeReservationTradeConsumeNoticeAPIRequest {
	return poolTaobaoLifeReservationTradeConsumeNoticeAPIRequest.Get().(*TaobaoLifeReservationTradeConsumeNoticeAPIRequest)
}

// ReleaseTaobaoLifeReservationTradeConsumeNoticeAPIRequest 将 TaobaoLifeReservationTradeConsumeNoticeAPIRequest 放入 sync.Pool
func ReleaseTaobaoLifeReservationTradeConsumeNoticeAPIRequest(v *TaobaoLifeReservationTradeConsumeNoticeAPIRequest) {
	v.Reset()
	poolTaobaoLifeReservationTradeConsumeNoticeAPIRequest.Put(v)
}

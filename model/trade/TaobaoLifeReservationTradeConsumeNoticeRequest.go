package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约外部核销 API请求
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销
*/
type TaobaoLifeReservationTradeConsumeNoticeRequest struct {
    model.Params
    // 淘宝主单号
    _tradeNo   string
    // 凭证ID
    _ticketId   string
}

// 初始化TaobaoLifeReservationTradeConsumeNoticeRequest对象
func NewTaobaoLifeReservationTradeConsumeNoticeRequest() *TaobaoLifeReservationTradeConsumeNoticeRequest{
    return &TaobaoLifeReservationTradeConsumeNoticeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetApiMethodName() string {
    return "taobao.life.reservation.trade.consume.notice"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 淘宝主单号
func (r *TaobaoLifeReservationTradeConsumeNoticeRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetTradeNo() string {
    return r._tradeNo
}
// TicketId Setter
// 凭证ID
func (r *TaobaoLifeReservationTradeConsumeNoticeRequest) SetTicketId(_ticketId string) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetTicketId() string {
    return r._ticketId
}

package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约外部核销 APIRequest
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销
*/
type TaobaoLifeReservationTradeConsumeNoticeRequest struct {
    model.Params

    // 淘宝主单号
    tradeNo   string 

    // 凭证ID
    ticketId   string 

}

func NewTaobaoLifeReservationTradeConsumeNoticeRequest() *TaobaoLifeReservationTradeConsumeNoticeRequest{
    return &TaobaoLifeReservationTradeConsumeNoticeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetApiMethodName() string {
    return "taobao.life.reservation.trade.consume.notice"
}

func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLifeReservationTradeConsumeNoticeRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *TaobaoLifeReservationTradeConsumeNoticeRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

func (r TaobaoLifeReservationTradeConsumeNoticeRequest) GetTicketId() string {
    return r.ticketId
}


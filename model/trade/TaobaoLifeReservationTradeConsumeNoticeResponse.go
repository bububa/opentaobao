package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约外部核销 APIResponse
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销
*/
type TaobaoLifeReservationTradeConsumeNoticeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLifeReservationTradeConsumeNoticeResponse `json:"taobao_life_reservation_trade_consume_notice_response,omitempty"`
}

type TaobaoLifeReservationTradeConsumeNoticeResponse struct {

    // 接口返回model
    Result   *LifeReservationTradeConsumeNoticeResult `json:"result,omitempty"`

}

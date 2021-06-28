package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约外部核销 APIResponse
taobao.life.reservation.trade.consume.notice

生服团购商品，购后预约。外部ISV进行核销
*/
type TaobaoLifeReservationTradeConsumeNoticeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"life_reservation_trade_consume_notice_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *LifeReservationTradeConsumeNoticeResult `json:"result,omitempty" xml:"
package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaolifereservationtradeconsumenoticeAPIResponse 生服购后预约外部核销 API返回值
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
type TaobaolifereservationtradeconsumenoticeAPIResponse struct {
	model.CommonResponse
	TaobaolifereservationtradeconsumenoticeAPIResponseModel
}

// TaobaolifereservationtradeconsumenoticeAPIResponseModel is 生服购后预约外部核销 成功返回结果
type TaobaolifereservationtradeconsumenoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"life_reservation_trade_consume_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaolifereservationtradeconsumenoticeResult `json:"result,omitempty" xml:"result,omitempty"`
}

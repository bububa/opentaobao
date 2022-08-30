package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketOrderVerifyAPIResponse 订单核销通知 API返回值
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
type AlitripTravelHotelticketOrderVerifyAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketOrderVerifyAPIResponseModel
}

// AlitripTravelHotelticketOrderVerifyAPIResponseModel is 订单核销通知 成功返回结果
type AlitripTravelHotelticketOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

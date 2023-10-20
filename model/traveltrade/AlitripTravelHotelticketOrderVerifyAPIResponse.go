package traveltrade

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelHotelticketOrderVerifyAPIResponseModel).Reset()
}

// AlitripTravelHotelticketOrderVerifyAPIResponseModel is 订单核销通知 成功返回结果
type AlitripTravelHotelticketOrderVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTravelHotelticketOrderVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelHotelticketOrderVerifyAPIResponse)
	},
}

// GetAlitripTravelHotelticketOrderVerifyAPIResponse 从 sync.Pool 获取 AlitripTravelHotelticketOrderVerifyAPIResponse
func GetAlitripTravelHotelticketOrderVerifyAPIResponse() *AlitripTravelHotelticketOrderVerifyAPIResponse {
	return poolAlitripTravelHotelticketOrderVerifyAPIResponse.Get().(*AlitripTravelHotelticketOrderVerifyAPIResponse)
}

// ReleaseAlitripTravelHotelticketOrderVerifyAPIResponse 将 AlitripTravelHotelticketOrderVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelHotelticketOrderVerifyAPIResponse(v *AlitripTravelHotelticketOrderVerifyAPIResponse) {
	v.Reset()
	poolAlitripTravelHotelticketOrderVerifyAPIResponse.Put(v)
}

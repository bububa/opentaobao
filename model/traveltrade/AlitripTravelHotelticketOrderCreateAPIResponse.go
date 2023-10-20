package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketOrderCreateAPIResponse 创单(支付订单)通知 API返回值
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
type AlitripTravelHotelticketOrderCreateAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelHotelticketOrderCreateAPIResponseModel).Reset()
}

// AlitripTravelHotelticketOrderCreateAPIResponseModel is 创单(支付订单)通知 成功返回结果
type AlitripTravelHotelticketOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTravelHotelticketOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelHotelticketOrderCreateAPIResponse)
	},
}

// GetAlitripTravelHotelticketOrderCreateAPIResponse 从 sync.Pool 获取 AlitripTravelHotelticketOrderCreateAPIResponse
func GetAlitripTravelHotelticketOrderCreateAPIResponse() *AlitripTravelHotelticketOrderCreateAPIResponse {
	return poolAlitripTravelHotelticketOrderCreateAPIResponse.Get().(*AlitripTravelHotelticketOrderCreateAPIResponse)
}

// ReleaseAlitripTravelHotelticketOrderCreateAPIResponse 将 AlitripTravelHotelticketOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelHotelticketOrderCreateAPIResponse(v *AlitripTravelHotelticketOrderCreateAPIResponse) {
	v.Reset()
	poolAlitripTravelHotelticketOrderCreateAPIResponse.Put(v)
}

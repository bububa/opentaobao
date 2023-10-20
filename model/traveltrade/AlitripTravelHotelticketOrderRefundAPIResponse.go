package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketOrderRefundAPIResponse 退款结结果通知 API返回值
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
type AlitripTravelHotelticketOrderRefundAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketOrderRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelHotelticketOrderRefundAPIResponseModel).Reset()
}

// AlitripTravelHotelticketOrderRefundAPIResponseModel is 退款结结果通知 成功返回结果
type AlitripTravelHotelticketOrderRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_order_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketOrderRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTravelHotelticketOrderRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelHotelticketOrderRefundAPIResponse)
	},
}

// GetAlitripTravelHotelticketOrderRefundAPIResponse 从 sync.Pool 获取 AlitripTravelHotelticketOrderRefundAPIResponse
func GetAlitripTravelHotelticketOrderRefundAPIResponse() *AlitripTravelHotelticketOrderRefundAPIResponse {
	return poolAlitripTravelHotelticketOrderRefundAPIResponse.Get().(*AlitripTravelHotelticketOrderRefundAPIResponse)
}

// ReleaseAlitripTravelHotelticketOrderRefundAPIResponse 将 AlitripTravelHotelticketOrderRefundAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelHotelticketOrderRefundAPIResponse(v *AlitripTravelHotelticketOrderRefundAPIResponse) {
	v.Reset()
	poolAlitripTravelHotelticketOrderRefundAPIResponse.Put(v)
}

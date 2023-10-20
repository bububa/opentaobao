package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketProductProductupdatepushAPIResponse 产品批量变更推送通知 API返回值
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
type AlitripTravelHotelticketProductProductupdatepushAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketProductProductupdatepushAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketProductProductupdatepushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelHotelticketProductProductupdatepushAPIResponseModel).Reset()
}

// AlitripTravelHotelticketProductProductupdatepushAPIResponseModel is 产品批量变更推送通知 成功返回结果
type AlitripTravelHotelticketProductProductupdatepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_product_productupdatepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketProductProductupdatepushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTravelHotelticketProductProductupdatepushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelHotelticketProductProductupdatepushAPIResponse)
	},
}

// GetAlitripTravelHotelticketProductProductupdatepushAPIResponse 从 sync.Pool 获取 AlitripTravelHotelticketProductProductupdatepushAPIResponse
func GetAlitripTravelHotelticketProductProductupdatepushAPIResponse() *AlitripTravelHotelticketProductProductupdatepushAPIResponse {
	return poolAlitripTravelHotelticketProductProductupdatepushAPIResponse.Get().(*AlitripTravelHotelticketProductProductupdatepushAPIResponse)
}

// ReleaseAlitripTravelHotelticketProductProductupdatepushAPIResponse 将 AlitripTravelHotelticketProductProductupdatepushAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelHotelticketProductProductupdatepushAPIResponse(v *AlitripTravelHotelticketProductProductupdatepushAPIResponse) {
	v.Reset()
	poolAlitripTravelHotelticketProductProductupdatepushAPIResponse.Put(v)
}

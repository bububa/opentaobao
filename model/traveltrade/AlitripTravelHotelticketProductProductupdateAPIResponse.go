package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketProductProductupdateAPIResponse 产品批量变更通知 API返回值
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
type AlitripTravelHotelticketProductProductupdateAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketProductProductupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketProductProductupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelHotelticketProductProductupdateAPIResponseModel).Reset()
}

// AlitripTravelHotelticketProductProductupdateAPIResponseModel is 产品批量变更通知 成功返回结果
type AlitripTravelHotelticketProductProductupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_product_productupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelHotelticketProductProductupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlitripTravelHotelticketProductProductupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelHotelticketProductProductupdateAPIResponse)
	},
}

// GetAlitripTravelHotelticketProductProductupdateAPIResponse 从 sync.Pool 获取 AlitripTravelHotelticketProductProductupdateAPIResponse
func GetAlitripTravelHotelticketProductProductupdateAPIResponse() *AlitripTravelHotelticketProductProductupdateAPIResponse {
	return poolAlitripTravelHotelticketProductProductupdateAPIResponse.Get().(*AlitripTravelHotelticketProductProductupdateAPIResponse)
}

// ReleaseAlitripTravelHotelticketProductProductupdateAPIResponse 将 AlitripTravelHotelticketProductProductupdateAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelHotelticketProductProductupdateAPIResponse(v *AlitripTravelHotelticketProductProductupdateAPIResponse) {
	v.Reset()
	poolAlitripTravelHotelticketProductProductupdateAPIResponse.Put(v)
}

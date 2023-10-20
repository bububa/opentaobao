package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchLowPriceAPIResponse 商旅酒店api分销-酒店最低价 API返回值
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
type AlitripBtripHotelDistributionSearchLowPriceAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionSearchLowPriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchLowPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionSearchLowPriceAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionSearchLowPriceAPIResponseModel is 商旅酒店api分销-酒店最低价 成功返回结果
type AlitripBtripHotelDistributionSearchLowPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_low_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchLowPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionSearchLowPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionSearchLowPriceAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionSearchLowPriceAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchLowPriceAPIResponse
func GetAlitripBtripHotelDistributionSearchLowPriceAPIResponse() *AlitripBtripHotelDistributionSearchLowPriceAPIResponse {
	return poolAlitripBtripHotelDistributionSearchLowPriceAPIResponse.Get().(*AlitripBtripHotelDistributionSearchLowPriceAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionSearchLowPriceAPIResponse 将 AlitripBtripHotelDistributionSearchLowPriceAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchLowPriceAPIResponse(v *AlitripBtripHotelDistributionSearchLowPriceAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchLowPriceAPIResponse.Put(v)
}

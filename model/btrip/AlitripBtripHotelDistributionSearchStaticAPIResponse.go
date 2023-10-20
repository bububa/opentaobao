package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchStaticAPIResponse 商旅酒店api分销-酒店静态信息接口 API返回值
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
type AlitripBtripHotelDistributionSearchStaticAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionSearchStaticAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchStaticAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionSearchStaticAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionSearchStaticAPIResponseModel is 商旅酒店api分销-酒店静态信息接口 成功返回结果
type AlitripBtripHotelDistributionSearchStaticAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_static_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchStaticAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionSearchStaticAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionSearchStaticAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionSearchStaticAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchStaticAPIResponse
func GetAlitripBtripHotelDistributionSearchStaticAPIResponse() *AlitripBtripHotelDistributionSearchStaticAPIResponse {
	return poolAlitripBtripHotelDistributionSearchStaticAPIResponse.Get().(*AlitripBtripHotelDistributionSearchStaticAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionSearchStaticAPIResponse 将 AlitripBtripHotelDistributionSearchStaticAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchStaticAPIResponse(v *AlitripBtripHotelDistributionSearchStaticAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchStaticAPIResponse.Put(v)
}

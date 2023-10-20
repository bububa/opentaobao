package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchDetailAPIResponse 商旅酒店api分销-详情报价接口 API返回值
// alitrip.btrip.hotel.distribution.search.detail
//
// 商旅酒店api分销-详情报价接口
type AlitripBtripHotelDistributionSearchDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionSearchDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripHotelDistributionSearchDetailAPIResponseModel).Reset()
}

// AlitripBtripHotelDistributionSearchDetailAPIResponseModel is 商旅酒店api分销-详情报价接口 成功返回结果
type AlitripBtripHotelDistributionSearchDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 详情报价回参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripHotelDistributionSearchDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripHotelDistributionSearchDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripHotelDistributionSearchDetailAPIResponse)
	},
}

// GetAlitripBtripHotelDistributionSearchDetailAPIResponse 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchDetailAPIResponse
func GetAlitripBtripHotelDistributionSearchDetailAPIResponse() *AlitripBtripHotelDistributionSearchDetailAPIResponse {
	return poolAlitripBtripHotelDistributionSearchDetailAPIResponse.Get().(*AlitripBtripHotelDistributionSearchDetailAPIResponse)
}

// ReleaseAlitripBtripHotelDistributionSearchDetailAPIResponse 将 AlitripBtripHotelDistributionSearchDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchDetailAPIResponse(v *AlitripBtripHotelDistributionSearchDetailAPIResponse) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchDetailAPIResponse.Put(v)
}

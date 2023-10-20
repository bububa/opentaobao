package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightIndustrySearchAPIResponse 机票行业搜索接口 API返回值
// alitrip.btrip.supplychain.flight.industry.search
//
// 【商旅】机票行业搜索
type AlitripBtripSupplychainFlightIndustrySearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainFlightIndustrySearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightIndustrySearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainFlightIndustrySearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainFlightIndustrySearchAPIResponseModel is 机票行业搜索接口 成功返回结果
type AlitripBtripSupplychainFlightIndustrySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_industry_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainFlightIndustrySearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainFlightIndustrySearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainFlightIndustrySearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainFlightIndustrySearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainFlightIndustrySearchAPIResponse
func GetAlitripBtripSupplychainFlightIndustrySearchAPIResponse() *AlitripBtripSupplychainFlightIndustrySearchAPIResponse {
	return poolAlitripBtripSupplychainFlightIndustrySearchAPIResponse.Get().(*AlitripBtripSupplychainFlightIndustrySearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainFlightIndustrySearchAPIResponse 将 AlitripBtripSupplychainFlightIndustrySearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainFlightIndustrySearchAPIResponse(v *AlitripBtripSupplychainFlightIndustrySearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainFlightIndustrySearchAPIResponse.Put(v)
}

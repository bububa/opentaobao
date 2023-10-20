package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainBusIndustrySearchAPIResponse 汽车票行业搜索接口 API返回值
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
type AlitripBtripSupplychainBusIndustrySearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainBusIndustrySearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainBusIndustrySearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripSupplychainBusIndustrySearchAPIResponseModel).Reset()
}

// AlitripBtripSupplychainBusIndustrySearchAPIResponseModel is 汽车票行业搜索接口 成功返回结果
type AlitripBtripSupplychainBusIndustrySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_bus_industry_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripSupplychainBusIndustrySearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripSupplychainBusIndustrySearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripSupplychainBusIndustrySearchAPIResponse)
	},
}

// GetAlitripBtripSupplychainBusIndustrySearchAPIResponse 从 sync.Pool 获取 AlitripBtripSupplychainBusIndustrySearchAPIResponse
func GetAlitripBtripSupplychainBusIndustrySearchAPIResponse() *AlitripBtripSupplychainBusIndustrySearchAPIResponse {
	return poolAlitripBtripSupplychainBusIndustrySearchAPIResponse.Get().(*AlitripBtripSupplychainBusIndustrySearchAPIResponse)
}

// ReleaseAlitripBtripSupplychainBusIndustrySearchAPIResponse 将 AlitripBtripSupplychainBusIndustrySearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripSupplychainBusIndustrySearchAPIResponse(v *AlitripBtripSupplychainBusIndustrySearchAPIResponse) {
	v.Reset()
	poolAlitripBtripSupplychainBusIndustrySearchAPIResponse.Put(v)
}

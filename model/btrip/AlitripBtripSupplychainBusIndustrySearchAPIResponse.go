package btrip

import (
	"encoding/xml"

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

// AlitripBtripSupplychainBusIndustrySearchAPIResponseModel is 汽车票行业搜索接口 成功返回结果
type AlitripBtripSupplychainBusIndustrySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_bus_industry_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票行业搜索接口 APIResponse
alitrip.btrip.supplychain.bus.industry.search

汽车票行业搜索接口
*/
type AlitripBtripSupplychainBusIndustrySearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainBusIndustrySearchResponse
}

type AlitripBtripSupplychainBusIndustrySearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_bus_industry_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

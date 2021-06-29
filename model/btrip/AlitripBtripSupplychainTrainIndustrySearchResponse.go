package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票行业搜索接口 APIResponse
alitrip.btrip.supplychain.train.industry.search

【商旅】火车票行业搜索接口
*/
type AlitripBtripSupplychainTrainIndustrySearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainTrainIndustrySearchResponse
}

type AlitripBtripSupplychainTrainIndustrySearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_train_industry_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

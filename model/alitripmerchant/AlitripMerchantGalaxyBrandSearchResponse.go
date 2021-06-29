package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-品牌搜索 APIResponse
alitrip.merchant.galaxy.brand.search

星河服务=获取雅高品牌信息
*/
type AlitripMerchantGalaxyBrandSearchAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyBrandSearchResponse
}

type AlitripMerchantGalaxyBrandSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_brand_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlitripMerchantGalaxyBrandSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

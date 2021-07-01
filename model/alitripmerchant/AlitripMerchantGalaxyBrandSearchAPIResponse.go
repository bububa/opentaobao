package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-品牌搜索 API返回值 
alitrip.merchant.galaxy.brand.search

星河服务=获取雅高品牌信息
*/
type AlitripMerchantGalaxyBrandSearchAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyBrandSearchAPIResponseModel
}

// 星河-品牌搜索 成功返回结果
type AlitripMerchantGalaxyBrandSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_brand_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlitripMerchantGalaxyBrandSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

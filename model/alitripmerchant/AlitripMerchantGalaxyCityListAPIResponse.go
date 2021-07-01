package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市列表展示 API返回值 
alitrip.merchant.galaxy.city.list

雅高酒店城市列表展示，并且首字母列出酒店城市
*/
type AlitripMerchantGalaxyCityListAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyCityListAPIResponseModel
}

// 星河-酒店城市列表展示 成功返回结果
type AlitripMerchantGalaxyCityListAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_city_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *AlitripMerchantGalaxyCityListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

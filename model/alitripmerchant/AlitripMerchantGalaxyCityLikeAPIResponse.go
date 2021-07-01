package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市模糊查询 API返回值 
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息
*/
type AlitripMerchantGalaxyCityLikeAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyCityLikeAPIResponseModel
}

// 星河-酒店城市模糊查询 成功返回结果
type AlitripMerchantGalaxyCityLikeAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_city_like_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *AlitripMerchantGalaxyCityLikeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

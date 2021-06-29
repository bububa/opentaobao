package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市模糊查询 APIResponse
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息
*/
type AlitripMerchantGalaxyCityLikeAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyCityLikeResponse
}

type AlitripMerchantGalaxyCityLikeResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_city_like_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}

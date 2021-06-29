package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取国家列表 APIResponse
aliexpress.social.country.get

获取目前AE支持的国家列表
*/
type AliexpressSocialCountryGetAPIResponse struct {
    model.CommonResponse
    AliexpressSocialCountryGetResponse
}

type AliexpressSocialCountryGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_country_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ItemPickPagingResult
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

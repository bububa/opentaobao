package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取国家列表 APIResponse
aliexpress.social.country.get

获取目前AE支持的国家列表
*/
type AliexpressSocialCountryGetAPIResponse struct {
    model.CommonResponse
    Response *AliexpressSocialCountryGetResponse `json:"aliexpress_social_country_get_response,omitempty"`
}

type AliexpressSocialCountryGetResponse struct {

    // ItemPickPagingResult
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}

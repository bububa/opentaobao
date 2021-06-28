package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
Locale获取接口 APIResponse
aliexpress.social.locale.get

新增Locale获取接口
*/
type AliexpressSocialLocaleGetAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressSocialLocaleGetResponse `json:"aliexpress_social_locale_get_response,omitempty"` 
    AliexpressSocialLocaleGetResponse
}

/* model for simplify = false
type AliexpressSocialLocaleGetResponse struct {

    // 包类型
    
    Result  *struct {
        ItemPickPagingResult  *ItemPickPagingResult `json:"item_pick_paging_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AliexpressSocialLocaleGetResponse struct {

    // 包类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}

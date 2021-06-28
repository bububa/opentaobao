package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Locale获取接口 APIResponse
aliexpress.social.locale.get

新增Locale获取接口
*/
type AliexpressSocialLocaleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aliexpress_social_locale_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 包类型
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"
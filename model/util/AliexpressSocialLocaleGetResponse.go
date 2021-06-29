package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Locale获取接口 API返回值 
aliexpress.social.locale.get

新增Locale获取接口
*/
type AliexpressSocialLocaleGetAPIResponse struct {
    model.CommonResponse
    AliexpressSocialLocaleGetResponse
}

// Locale获取接口 成功返回结果
type AliexpressSocialLocaleGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_locale_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 包类型
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

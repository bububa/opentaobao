package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
币种获取接口 APIResponse
aliexpress.social.currency.get

获取目前AE社交支持的币种
*/
type AliexpressSocialCurrencyGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aliexpress_social_currency_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 包类型
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"
package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
币种获取接口 APIResponse
aliexpress.social.currency.get

获取目前AE社交支持的币种
*/
type AliexpressSocialCurrencyGetAPIResponse struct {
    model.CommonResponse
    Response *AliexpressSocialCurrencyGetResponse `json:"aliexpress_social_currency_get_response,omitempty"`
}

type AliexpressSocialCurrencyGetResponse struct {

    // 包类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}

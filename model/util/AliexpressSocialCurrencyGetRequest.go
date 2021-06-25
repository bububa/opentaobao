package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
币种获取接口 APIRequest
aliexpress.social.currency.get

获取目前AE社交支持的币种
*/
type AliexpressSocialCurrencyGetRequest struct {
    model.Params

}

func NewAliexpressSocialCurrencyGetRequest() *AliexpressSocialCurrencyGetRequest{
    return &AliexpressSocialCurrencyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialCurrencyGetRequest) GetApiMethodName() string {
    return "aliexpress.social.currency.get"
}

func (r AliexpressSocialCurrencyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}



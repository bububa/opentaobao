package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
币种获取接口 API请求
aliexpress.social.currency.get

获取目前AE社交支持的币种
*/
type AliexpressSocialCurrencyGetRequest struct {
    model.Params
}

// 初始化AliexpressSocialCurrencyGetRequest对象
func NewAliexpressSocialCurrencyGetRequest() *AliexpressSocialCurrencyGetRequest{
    return &AliexpressSocialCurrencyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialCurrencyGetRequest) GetApiMethodName() string {
    return "aliexpress.social.currency.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialCurrencyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

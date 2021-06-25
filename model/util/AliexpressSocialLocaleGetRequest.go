package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Locale获取接口 APIRequest
aliexpress.social.locale.get

新增Locale获取接口
*/
type AliexpressSocialLocaleGetRequest struct {
    model.Params

}

func NewAliexpressSocialLocaleGetRequest() *AliexpressSocialLocaleGetRequest{
    return &AliexpressSocialLocaleGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialLocaleGetRequest) GetApiMethodName() string {
    return "aliexpress.social.locale.get"
}

func (r AliexpressSocialLocaleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}



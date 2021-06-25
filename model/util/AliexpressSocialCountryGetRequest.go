package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取国家列表 APIRequest
aliexpress.social.country.get

获取目前AE支持的国家列表
*/
type AliexpressSocialCountryGetRequest struct {
    model.Params

    // 语言
    language   string 

}

func NewAliexpressSocialCountryGetRequest() *AliexpressSocialCountryGetRequest{
    return &AliexpressSocialCountryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialCountryGetRequest) GetApiMethodName() string {
    return "aliexpress.social.country.get"
}

func (r AliexpressSocialCountryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialCountryGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AliexpressSocialCountryGetRequest) GetLanguage() string {
    return r.language
}


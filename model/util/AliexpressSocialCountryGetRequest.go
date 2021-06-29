package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取国家列表 API请求
aliexpress.social.country.get

获取目前AE支持的国家列表
*/
type AliexpressSocialCountryGetRequest struct {
    model.Params
    // 语言
    _language   string
}

// 初始化AliexpressSocialCountryGetRequest对象
func NewAliexpressSocialCountryGetRequest() *AliexpressSocialCountryGetRequest{
    return &AliexpressSocialCountryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialCountryGetRequest) GetApiMethodName() string {
    return "aliexpress.social.country.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialCountryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Language Setter
// 语言
func (r *AliexpressSocialCountryGetRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AliexpressSocialCountryGetRequest) GetLanguage() string {
    return r._language
}

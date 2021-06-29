package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Locale获取接口 API请求
aliexpress.social.locale.get

新增Locale获取接口
*/
type AliexpressSocialLocaleGetRequest struct {
    model.Params
}

// 初始化AliexpressSocialLocaleGetRequest对象
func NewAliexpressSocialLocaleGetRequest() *AliexpressSocialLocaleGetRequest{
    return &AliexpressSocialLocaleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialLocaleGetRequest) GetApiMethodName() string {
    return "aliexpress.social.locale.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialLocaleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

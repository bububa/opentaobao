package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 API请求
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetRequest struct {
    model.Params
    // Locale值，格式为language+"_"+country
    locale   string
}

// 初始化AliexpressSocialDiscategoryGetRequest对象
func NewAliexpressSocialDiscategoryGetRequest() *AliexpressSocialDiscategoryGetRequest{
    return &AliexpressSocialDiscategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialDiscategoryGetRequest) GetApiMethodName() string {
    return "aliexpress.social.discategory.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialDiscategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Locale Setter
// Locale值，格式为language+"_"+country
func (r *AliexpressSocialDiscategoryGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r AliexpressSocialDiscategoryGetRequest) GetLocale() string {
    return r.locale
}

package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 APIRequest
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetRequest struct {
    model.Params

    // Locale值，格式为language+"_"+country
    locale   string 

}

func NewAliexpressSocialDiscategoryGetRequest() *AliexpressSocialDiscategoryGetRequest{
    return &AliexpressSocialDiscategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialDiscategoryGetRequest) GetApiMethodName() string {
    return "aliexpress.social.discategory.get"
}

func (r AliexpressSocialDiscategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialDiscategoryGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r AliexpressSocialDiscategoryGetRequest) GetLocale() string {
    return r.locale
}


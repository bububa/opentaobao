package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量推广类目信息获取API APIRequest
aliexpress.affiliate.category.get

获取AE流量推广类目的API
*/
type AliexpressAffiliateCategoryGetRequest struct {
    model.Params

    // 请求安全签名
    appSignature   string 

}

func NewAliexpressAffiliateCategoryGetRequest() *AliexpressAffiliateCategoryGetRequest{
    return &AliexpressAffiliateCategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateCategoryGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.category.get"
}

func (r AliexpressAffiliateCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateCategoryGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateCategoryGetRequest) GetAppSignature() string {
    return r.appSignature
}


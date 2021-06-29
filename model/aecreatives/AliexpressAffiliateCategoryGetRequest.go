package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量推广类目信息获取API API请求
aliexpress.affiliate.category.get

获取AE流量推广类目的API
*/
type AliexpressAffiliateCategoryGetRequest struct {
    model.Params
    // 请求安全签名
    _appSignature   string
}

// 初始化AliexpressAffiliateCategoryGetRequest对象
func NewAliexpressAffiliateCategoryGetRequest() *AliexpressAffiliateCategoryGetRequest{
    return &AliexpressAffiliateCategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateCategoryGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.category.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 请求安全签名
func (r *AliexpressAffiliateCategoryGetRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateCategoryGetRequest) GetAppSignature() string {
    return r._appSignature
}

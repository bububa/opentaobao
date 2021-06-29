package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟主题推广活动信息获取 API请求
aliexpress.affiliate.featuredpromo.get

获取联盟主题推广活动信息
*/
type AliexpressAffiliateFeaturedpromoGetRequest struct {
    model.Params
    // 请求签名
    appSignature   string
    // 返回字段列表
    fields   string
}

// 初始化AliexpressAffiliateFeaturedpromoGetRequest对象
func NewAliexpressAffiliateFeaturedpromoGetRequest() *AliexpressAffiliateFeaturedpromoGetRequest{
    return &AliexpressAffiliateFeaturedpromoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateFeaturedpromoGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.featuredpromo.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateFeaturedpromoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateFeaturedpromoGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateFeaturedpromoGetRequest) GetAppSignature() string {
    return r.appSignature
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateFeaturedpromoGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateFeaturedpromoGetRequest) GetFields() string {
    return r.fields
}

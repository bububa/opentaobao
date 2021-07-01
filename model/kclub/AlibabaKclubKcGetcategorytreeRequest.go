package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询租户下类目树 API请求
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
*/
type AlibabaKclubKcGetcategorytreeAPIRequest struct {
    model.Params
    // 租户id
    _tenantId   int64
    // 鉴权参数
    _auth   *TenancyAuth
}

// 初始化AlibabaKclubKcGetcategorytreeAPIRequest对象
func NewAlibabaKclubKcGetcategorytreeRequest() *AlibabaKclubKcGetcategorytreeAPIRequest{
    return &AlibabaKclubKcGetcategorytreeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.getcategorytree"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantId Setter
// 租户id
func (r *AlibabaKclubKcGetcategorytreeAPIRequest) SetTenantId(_tenantId int64) error {
    r._tenantId = _tenantId
    r.Set("tenant_id", _tenantId)
    return nil
}

// TenantId Getter
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetTenantId() int64 {
    return r._tenantId
}
// Auth Setter
// 鉴权参数
func (r *AlibabaKclubKcGetcategorytreeAPIRequest) SetAuth(_auth *TenancyAuth) error {
    r._auth = _auth
    r.Set("auth", _auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetAuth() *TenancyAuth {
    return r._auth
}

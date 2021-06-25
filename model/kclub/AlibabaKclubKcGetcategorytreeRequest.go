package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询租户下类目树 APIRequest
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
*/
type AlibabaKclubKcGetcategorytreeRequest struct {
    model.Params

    // 租户id
    tenantId   int64 

    // 鉴权参数
    auth   *TenancyAuth 

}

func NewAlibabaKclubKcGetcategorytreeRequest() *AlibabaKclubKcGetcategorytreeRequest{
    return &AlibabaKclubKcGetcategorytreeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKclubKcGetcategorytreeRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.getcategorytree"
}

func (r AlibabaKclubKcGetcategorytreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKclubKcGetcategorytreeRequest) SetTenantId(tenantId int64) error {
    r.tenantId = tenantId
    r.Set("tenant_id", tenantId)
    return nil
}

func (r AlibabaKclubKcGetcategorytreeRequest) GetTenantId() int64 {
    return r.tenantId
}

func (r *AlibabaKclubKcGetcategorytreeRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

func (r AlibabaKclubKcGetcategorytreeRequest) GetAuth() *TenancyAuth {
    return r.auth
}


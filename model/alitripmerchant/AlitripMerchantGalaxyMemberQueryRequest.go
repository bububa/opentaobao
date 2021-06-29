package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取登录用户的信息 API请求
alitrip.merchant.galaxy.member.query

获取登录用户的信息
*/
type AlitripMerchantGalaxyMemberQueryRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // toekn
    _token   string
}

// 初始化AlitripMerchantGalaxyMemberQueryRequest对象
func NewAlitripMerchantGalaxyMemberQueryRequest() *AlitripMerchantGalaxyMemberQueryRequest{
    return &AlitripMerchantGalaxyMemberQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberQueryRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberQueryRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// toekn
func (r *AlitripMerchantGalaxyMemberQueryRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyMemberQueryRequest) GetToken() string {
    return r._token
}
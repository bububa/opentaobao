package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberQueryAPIRequest 星河-获取登录用户的信息 API请求
// alitrip.merchant.galaxy.member.query
//
// 获取登录用户的信息
type AlitripMerchantGalaxyMemberQueryAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// toekn
	_token string
}

// NewAlitripMerchantGalaxyMemberQueryRequest 初始化AlitripMerchantGalaxyMemberQueryAPIRequest对象
func NewAlitripMerchantGalaxyMemberQueryRequest() *AlitripMerchantGalaxyMemberQueryAPIRequest {
	return &AlitripMerchantGalaxyMemberQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// toekn
func (r *AlitripMerchantGalaxyMemberQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetToken() string {
	return r._token
}

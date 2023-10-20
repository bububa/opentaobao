package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberTokenAPIRequest 星河-校验token API请求
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
type AlitripMerchantGalaxyMemberTokenAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 访问携带toekn
	_token string
}

// NewAlitripMerchantGalaxyMemberTokenRequest 初始化AlitripMerchantGalaxyMemberTokenAPIRequest对象
func NewAlitripMerchantGalaxyMemberTokenRequest() *AlitripMerchantGalaxyMemberTokenAPIRequest {
	return &AlitripMerchantGalaxyMemberTokenAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberTokenAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.token"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberTokenAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 访问携带toekn
func (r *AlitripMerchantGalaxyMemberTokenAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyMemberTokenAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberTokenRequest()
	},
}

// GetAlitripMerchantGalaxyMemberTokenRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberTokenAPIRequest
func GetAlitripMerchantGalaxyMemberTokenAPIRequest() *AlitripMerchantGalaxyMemberTokenAPIRequest {
	return poolAlitripMerchantGalaxyMemberTokenAPIRequest.Get().(*AlitripMerchantGalaxyMemberTokenAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberTokenAPIRequest 将 AlitripMerchantGalaxyMemberTokenAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberTokenAPIRequest(v *AlitripMerchantGalaxyMemberTokenAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberTokenAPIRequest.Put(v)
}

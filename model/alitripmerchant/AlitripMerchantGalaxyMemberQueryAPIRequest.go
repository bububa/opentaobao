package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripMerchantGalaxyMemberQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberQueryRequest()
	},
}

// GetAlitripMerchantGalaxyMemberQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberQueryAPIRequest
func GetAlitripMerchantGalaxyMemberQueryAPIRequest() *AlitripMerchantGalaxyMemberQueryAPIRequest {
	return poolAlitripMerchantGalaxyMemberQueryAPIRequest.Get().(*AlitripMerchantGalaxyMemberQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberQueryAPIRequest 将 AlitripMerchantGalaxyMemberQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberQueryAPIRequest(v *AlitripMerchantGalaxyMemberQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberQueryAPIRequest.Put(v)
}

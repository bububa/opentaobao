package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberProviderRegisterAPIRequest 对外提供会员注册服务 API请求
// alitrip.merchant.galaxy.member.provider.register
//
// 星河产品=对外提供注册雅高会员服务
type AlitripMerchantGalaxyMemberProviderRegisterAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 注册入参
	_providerMemberParam *ProviderMemberParam
}

// NewAlitripMerchantGalaxyMemberProviderRegisterRequest 初始化AlitripMerchantGalaxyMemberProviderRegisterAPIRequest对象
func NewAlitripMerchantGalaxyMemberProviderRegisterRequest() *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest {
	return &AlitripMerchantGalaxyMemberProviderRegisterAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) Reset() {
	r._tenantKey = ""
	r._providerMemberParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.provider.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetProviderMemberParam is ProviderMemberParam Setter
// 注册入参
func (r *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) SetProviderMemberParam(_providerMemberParam *ProviderMemberParam) error {
	r._providerMemberParam = _providerMemberParam
	r.Set("provider_member_param", _providerMemberParam)
	return nil
}

// GetProviderMemberParam ProviderMemberParam Getter
func (r AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) GetProviderMemberParam() *ProviderMemberParam {
	return r._providerMemberParam
}

var poolAlitripMerchantGalaxyMemberProviderRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberProviderRegisterRequest()
	},
}

// GetAlitripMerchantGalaxyMemberProviderRegisterRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberProviderRegisterAPIRequest
func GetAlitripMerchantGalaxyMemberProviderRegisterAPIRequest() *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest {
	return poolAlitripMerchantGalaxyMemberProviderRegisterAPIRequest.Get().(*AlitripMerchantGalaxyMemberProviderRegisterAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberProviderRegisterAPIRequest 将 AlitripMerchantGalaxyMemberProviderRegisterAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberProviderRegisterAPIRequest(v *AlitripMerchantGalaxyMemberProviderRegisterAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberProviderRegisterAPIRequest.Put(v)
}

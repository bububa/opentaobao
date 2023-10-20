package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberproviderregisterAPIRequest 对外提供会员注册服务 API请求
// alitrip.merchant.galaxy.member.provider.register
//
// 星河产品=对外提供注册雅高会员服务
type AlitripmerchantgalaxymemberproviderregisterAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 注册入参
	_providerMemberParam *ProviderMemberParam
}

// NewAlitripmerchantgalaxymemberproviderregisterRequest 初始化AlitripmerchantgalaxymemberproviderregisterAPIRequest对象
func NewAlitripmerchantgalaxymemberproviderregisterRequest() *AlitripmerchantgalaxymemberproviderregisterAPIRequest {
	return &AlitripmerchantgalaxymemberproviderregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberproviderregisterAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.provider.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberproviderregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberproviderregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxymemberproviderregisterAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberproviderregisterAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetProviderMemberParam is ProviderMemberParam Setter
// 注册入参
func (r *AlitripmerchantgalaxymemberproviderregisterAPIRequest) SetProviderMemberParam(_providerMemberParam *ProviderMemberParam) error {
	r._providerMemberParam = _providerMemberParam
	r.Set("provider_member_param", _providerMemberParam)
	return nil
}

// GetProviderMemberParam ProviderMemberParam Getter
func (r AlitripmerchantgalaxymemberproviderregisterAPIRequest) GetProviderMemberParam() *ProviderMemberParam {
	return r._providerMemberParam
}

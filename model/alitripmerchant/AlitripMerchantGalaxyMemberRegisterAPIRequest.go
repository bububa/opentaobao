package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberRegisterAPIRequest 星河-微信小程序会员注册 API请求
// alitrip.merchant.galaxy.member.register
//
// 星河产品=微信小程序注册雅高会员服务
type AlitripMerchantGalaxyMemberRegisterAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 注册入参
	_memberParam *MemberParam
}

// NewAlitripMerchantGalaxyMemberRegisterRequest 初始化AlitripMerchantGalaxyMemberRegisterAPIRequest对象
func NewAlitripMerchantGalaxyMemberRegisterRequest() *AlitripMerchantGalaxyMemberRegisterAPIRequest {
	return &AlitripMerchantGalaxyMemberRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberRegisterAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberRegisterAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberRegisterAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetMemberParam is MemberParam Setter
// 注册入参
func (r *AlitripMerchantGalaxyMemberRegisterAPIRequest) SetMemberParam(_memberParam *MemberParam) error {
	r._memberParam = _memberParam
	r.Set("member_param", _memberParam)
	return nil
}

// GetMemberParam MemberParam Getter
func (r AlitripMerchantGalaxyMemberRegisterAPIRequest) GetMemberParam() *MemberParam {
	return r._memberParam
}

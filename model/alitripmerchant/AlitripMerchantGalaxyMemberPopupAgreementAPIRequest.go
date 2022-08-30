package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberPopupAgreementAPIRequest 小程序唤起协议弹窗 API请求
// alitrip.merchant.galaxy.member.popup.agreement
//
// 用户进入小程序后，根据用户是否授权协议，判断是否唤起协议弹窗
type AlitripMerchantGalaxyMemberPopupAgreementAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 微信用户标识
	_code string
}

// NewAlitripMerchantGalaxyMemberPopupAgreementRequest 初始化AlitripMerchantGalaxyMemberPopupAgreementAPIRequest对象
func NewAlitripMerchantGalaxyMemberPopupAgreementRequest() *AlitripMerchantGalaxyMemberPopupAgreementAPIRequest {
	return &AlitripMerchantGalaxyMemberPopupAgreementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.popup.agreement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCode is Code Setter
// 微信用户标识
func (r *AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripMerchantGalaxyMemberPopupAgreementAPIRequest) GetCode() string {
	return r._code
}

package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitypopupqueryAPIRequest 星河-获取雅高小程序营销抽奖首页弹窗 API请求
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
type AlitripmerchantgalaxyactivitypopupqueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 用户来源
	_userSiteCode int64
}

// NewAlitripmerchantgalaxyactivitypopupqueryRequest 初始化AlitripmerchantgalaxyactivitypopupqueryAPIRequest对象
func NewAlitripmerchantgalaxyactivitypopupqueryRequest() *AlitripmerchantgalaxyactivitypopupqueryAPIRequest {
	return &AlitripmerchantgalaxyactivitypopupqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.popup.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyactivitypopupqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyactivitypopupqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetToken() string {
	return r._token
}

// SetUserSiteCode is UserSiteCode Setter
// 用户来源
func (r *AlitripmerchantgalaxyactivitypopupqueryAPIRequest) SetUserSiteCode(_userSiteCode int64) error {
	r._userSiteCode = _userSiteCode
	r.Set("user_site_code", _userSiteCode)
	return nil
}

// GetUserSiteCode UserSiteCode Getter
func (r AlitripmerchantgalaxyactivitypopupqueryAPIRequest) GetUserSiteCode() int64 {
	return r._userSiteCode
}

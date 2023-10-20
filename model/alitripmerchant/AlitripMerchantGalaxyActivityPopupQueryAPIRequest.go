package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityPopupQueryAPIRequest 星河-获取雅高小程序营销抽奖首页弹窗 API请求
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
type AlitripMerchantGalaxyActivityPopupQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 用户来源
	_userSiteCode int64
}

// NewAlitripMerchantGalaxyActivityPopupQueryRequest 初始化AlitripMerchantGalaxyActivityPopupQueryAPIRequest对象
func NewAlitripMerchantGalaxyActivityPopupQueryRequest() *AlitripMerchantGalaxyActivityPopupQueryAPIRequest {
	return &AlitripMerchantGalaxyActivityPopupQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityPopupQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._userSiteCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.popup.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyActivityPopupQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyActivityPopupQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetToken() string {
	return r._token
}

// SetUserSiteCode is UserSiteCode Setter
// 用户来源
func (r *AlitripMerchantGalaxyActivityPopupQueryAPIRequest) SetUserSiteCode(_userSiteCode int64) error {
	r._userSiteCode = _userSiteCode
	r.Set("user_site_code", _userSiteCode)
	return nil
}

// GetUserSiteCode UserSiteCode Getter
func (r AlitripMerchantGalaxyActivityPopupQueryAPIRequest) GetUserSiteCode() int64 {
	return r._userSiteCode
}

var poolAlitripMerchantGalaxyActivityPopupQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityPopupQueryRequest()
	},
}

// GetAlitripMerchantGalaxyActivityPopupQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityPopupQueryAPIRequest
func GetAlitripMerchantGalaxyActivityPopupQueryAPIRequest() *AlitripMerchantGalaxyActivityPopupQueryAPIRequest {
	return poolAlitripMerchantGalaxyActivityPopupQueryAPIRequest.Get().(*AlitripMerchantGalaxyActivityPopupQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityPopupQueryAPIRequest 将 AlitripMerchantGalaxyActivityPopupQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityPopupQueryAPIRequest(v *AlitripMerchantGalaxyActivityPopupQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityPopupQueryAPIRequest.Put(v)
}

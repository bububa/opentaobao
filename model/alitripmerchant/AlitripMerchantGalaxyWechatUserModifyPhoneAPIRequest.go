package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest DFC-ID用户换绑手机号 API请求
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
type AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 微信向程序登录参数
	_loginParam *LoginParam
}

// NewAlitripMerchantGalaxyWechatUserModifyPhoneRequest 初始化AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest对象
func NewAlitripMerchantGalaxyWechatUserModifyPhoneRequest() *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest {
	return &AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._loginParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.user.modify.phone"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetToken() string {
	return r._token
}

// SetLoginParam is LoginParam Setter
// 微信向程序登录参数
func (r *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}

var poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatUserModifyPhoneRequest()
	},
}

// GetAlitripMerchantGalaxyWechatUserModifyPhoneRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest
func GetAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest() *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest {
	return poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest.Get().(*AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest 将 AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest(v *AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest.Put(v)
}

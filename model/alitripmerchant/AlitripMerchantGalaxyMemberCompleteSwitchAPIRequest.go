package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymembercompleteswitchAPIRequest 会员切换模式接口 API请求
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
type AlitripmerchantgalaxymembercompleteswitchAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 验证码
	_verificationCode string
	// 场景来源，记录什么场景调用什么接口
	_sceneSource string
}

// NewAlitripmerchantgalaxymembercompleteswitchRequest 初始化AlitripmerchantgalaxymembercompleteswitchAPIRequest对象
func NewAlitripmerchantgalaxymembercompleteswitchRequest() *AlitripmerchantgalaxymembercompleteswitchAPIRequest {
	return &AlitripmerchantgalaxymembercompleteswitchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.complete.switch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxymembercompleteswitchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxymembercompleteswitchAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetToken() string {
	return r._token
}

// SetVerificationCode is VerificationCode Setter
// 验证码
func (r *AlitripmerchantgalaxymembercompleteswitchAPIRequest) SetVerificationCode(_verificationCode string) error {
	r._verificationCode = _verificationCode
	r.Set("verification_code", _verificationCode)
	return nil
}

// GetVerificationCode VerificationCode Getter
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetVerificationCode() string {
	return r._verificationCode
}

// SetSceneSource is SceneSource Setter
// 场景来源，记录什么场景调用什么接口
func (r *AlitripmerchantgalaxymembercompleteswitchAPIRequest) SetSceneSource(_sceneSource string) error {
	r._sceneSource = _sceneSource
	r.Set("scene_source", _sceneSource)
	return nil
}

// GetSceneSource SceneSource Getter
func (r AlitripmerchantgalaxymembercompleteswitchAPIRequest) GetSceneSource() string {
	return r._sceneSource
}

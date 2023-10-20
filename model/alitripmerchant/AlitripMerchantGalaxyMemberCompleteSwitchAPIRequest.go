package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest 会员切换模式接口 API请求
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
type AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest struct {
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

// NewAlitripMerchantGalaxyMemberCompleteSwitchRequest 初始化AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest对象
func NewAlitripMerchantGalaxyMemberCompleteSwitchRequest() *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest {
	return &AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._verificationCode = ""
	r._sceneSource = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.complete.switch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetToken() string {
	return r._token
}

// SetVerificationCode is VerificationCode Setter
// 验证码
func (r *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) SetVerificationCode(_verificationCode string) error {
	r._verificationCode = _verificationCode
	r.Set("verification_code", _verificationCode)
	return nil
}

// GetVerificationCode VerificationCode Getter
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetVerificationCode() string {
	return r._verificationCode
}

// SetSceneSource is SceneSource Setter
// 场景来源，记录什么场景调用什么接口
func (r *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) SetSceneSource(_sceneSource string) error {
	r._sceneSource = _sceneSource
	r.Set("scene_source", _sceneSource)
	return nil
}

// GetSceneSource SceneSource Getter
func (r AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) GetSceneSource() string {
	return r._sceneSource
}

var poolAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberCompleteSwitchRequest()
	},
}

// GetAlitripMerchantGalaxyMemberCompleteSwitchRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest
func GetAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest() *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest {
	return poolAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest.Get().(*AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest 将 AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest(v *AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberCompleteSwitchAPIRequest.Put(v)
}

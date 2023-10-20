package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberLoginDerbyAPIRequest 小程序通过德比登入（会员认证） API请求
// alitrip.merchant.galaxy.member.login.derby
//
// 会员认证(德比切换接口，包含授权接口)   返回认证及授权状态
type AlitripMerchantGalaxyMemberLoginDerbyAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 场景来源，记录什么场景调用什么接口
	_sceneSource string
	// 德比认证会员信息入参
	_derbyAuthenticationParam *DerbyAuthenticationParam
}

// NewAlitripMerchantGalaxyMemberLoginDerbyRequest 初始化AlitripMerchantGalaxyMemberLoginDerbyAPIRequest对象
func NewAlitripMerchantGalaxyMemberLoginDerbyRequest() *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest {
	return &AlitripMerchantGalaxyMemberLoginDerbyAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._sceneSource = ""
	r._derbyAuthenticationParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.login.derby"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetToken() string {
	return r._token
}

// SetSceneSource is SceneSource Setter
// 场景来源，记录什么场景调用什么接口
func (r *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) SetSceneSource(_sceneSource string) error {
	r._sceneSource = _sceneSource
	r.Set("scene_source", _sceneSource)
	return nil
}

// GetSceneSource SceneSource Getter
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetSceneSource() string {
	return r._sceneSource
}

// SetDerbyAuthenticationParam is DerbyAuthenticationParam Setter
// 德比认证会员信息入参
func (r *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) SetDerbyAuthenticationParam(_derbyAuthenticationParam *DerbyAuthenticationParam) error {
	r._derbyAuthenticationParam = _derbyAuthenticationParam
	r.Set("derby_authentication_param", _derbyAuthenticationParam)
	return nil
}

// GetDerbyAuthenticationParam DerbyAuthenticationParam Getter
func (r AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) GetDerbyAuthenticationParam() *DerbyAuthenticationParam {
	return r._derbyAuthenticationParam
}

var poolAlitripMerchantGalaxyMemberLoginDerbyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberLoginDerbyRequest()
	},
}

// GetAlitripMerchantGalaxyMemberLoginDerbyRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberLoginDerbyAPIRequest
func GetAlitripMerchantGalaxyMemberLoginDerbyAPIRequest() *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest {
	return poolAlitripMerchantGalaxyMemberLoginDerbyAPIRequest.Get().(*AlitripMerchantGalaxyMemberLoginDerbyAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberLoginDerbyAPIRequest 将 AlitripMerchantGalaxyMemberLoginDerbyAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberLoginDerbyAPIRequest(v *AlitripMerchantGalaxyMemberLoginDerbyAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberLoginDerbyAPIRequest.Put(v)
}

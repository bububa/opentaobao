package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberregisterderbyAPIRequest 会员注册(新版注册接口对接德比) API请求
// alitrip.merchant.galaxy.member.register.derby
//
// 会员注册(新版注册接口对接德比)，返回手机号/邮箱/注册状态
type AlitripmerchantgalaxymemberregisterderbyAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 场景来源，记录什么场景调用什么接口
	_sceneSource string
	// 用户注册信息
	_memberParam *MemberParam
}

// NewAlitripmerchantgalaxymemberregisterderbyRequest 初始化AlitripmerchantgalaxymemberregisterderbyAPIRequest对象
func NewAlitripmerchantgalaxymemberregisterderbyRequest() *AlitripmerchantgalaxymemberregisterderbyAPIRequest {
	return &AlitripmerchantgalaxymemberregisterderbyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.register.derby"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxymemberregisterderbyAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxymemberregisterderbyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetToken() string {
	return r._token
}

// SetSceneSource is SceneSource Setter
// 场景来源，记录什么场景调用什么接口
func (r *AlitripmerchantgalaxymemberregisterderbyAPIRequest) SetSceneSource(_sceneSource string) error {
	r._sceneSource = _sceneSource
	r.Set("scene_source", _sceneSource)
	return nil
}

// GetSceneSource SceneSource Getter
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetSceneSource() string {
	return r._sceneSource
}

// SetMemberParam is MemberParam Setter
// 用户注册信息
func (r *AlitripmerchantgalaxymemberregisterderbyAPIRequest) SetMemberParam(_memberParam *MemberParam) error {
	r._memberParam = _memberParam
	r.Set("member_param", _memberParam)
	return nil
}

// GetMemberParam MemberParam Getter
func (r AlitripmerchantgalaxymemberregisterderbyAPIRequest) GetMemberParam() *MemberParam {
	return r._memberParam
}

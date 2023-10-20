package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyuserriskAPIRequest 查询微信账号的风险等级 API请求
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
type AlitripmerchantgalaxyuserriskAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 用户ip
	_ip string
	// 邮箱
	_email string
	// 场景值
	_scene int64
}

// NewAlitripmerchantgalaxyuserriskRequest 初始化AlitripmerchantgalaxyuserriskAPIRequest对象
func NewAlitripmerchantgalaxyuserriskRequest() *AlitripmerchantgalaxyuserriskAPIRequest {
	return &AlitripmerchantgalaxyuserriskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.user.risk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyuserriskAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyuserriskAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetToken() string {
	return r._token
}

// SetIp is Ip Setter
// 用户ip
func (r *AlitripmerchantgalaxyuserriskAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetIp() string {
	return r._ip
}

// SetEmail is Email Setter
// 邮箱
func (r *AlitripmerchantgalaxyuserriskAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetEmail() string {
	return r._email
}

// SetScene is Scene Setter
// 场景值
func (r *AlitripmerchantgalaxyuserriskAPIRequest) SetScene(_scene int64) error {
	r._scene = _scene
	r.Set("scene", _scene)
	return nil
}

// GetScene Scene Getter
func (r AlitripmerchantgalaxyuserriskAPIRequest) GetScene() int64 {
	return r._scene
}

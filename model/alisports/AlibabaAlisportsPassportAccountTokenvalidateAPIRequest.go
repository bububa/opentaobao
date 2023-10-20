package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccounttokenvalidateAPIRequest 阿里体育会员系统帐号登录注册token验证接口 API请求
// alibaba.alisports.passport.account.tokenvalidate
//
// 阿里体育会员系统帐号登录注册token验证接口
type AlibabaalisportspassportaccounttokenvalidateAPIRequest struct {
	model.Params
	// 业务方appkey
	_alispAppKey string
	// 签名
	_alispSign string
	// token
	_token string
	// 时间戳
	_alispTime string
	// 一键登录参数
	_secret string
	// json字符串，传入扩展字段
	_extInfo string
	// 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
	_mtopAppkey string
	// 注册用户类型
	_userType int64
}

// NewAlibabaalisportspassportaccounttokenvalidateRequest 初始化AlibabaalisportspassportaccounttokenvalidateAPIRequest对象
func NewAlibabaalisportspassportaccounttokenvalidateRequest() *AlibabaalisportspassportaccounttokenvalidateAPIRequest {
	return &AlibabaalisportspassportaccounttokenvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.tokenvalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 业务方appkey
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispSign is AlispSign Setter
// 签名
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetToken is Token Setter
// token
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetToken() string {
	return r._token
}

// SetAlispTime is AlispTime Setter
// 时间戳
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetSecret is Secret Setter
// 一键登录参数
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetSecret(_secret string) error {
	r._secret = _secret
	r.Set("secret", _secret)
	return nil
}

// GetSecret Secret Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetSecret() string {
	return r._secret
}

// SetExtInfo is ExtInfo Setter
// json字符串，传入扩展字段
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetExtInfo() string {
	return r._extInfo
}

// SetMtopAppkey is MtopAppkey Setter
// 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetMtopAppkey(_mtopAppkey string) error {
	r._mtopAppkey = _mtopAppkey
	r.Set("mtop_appkey", _mtopAppkey)
	return nil
}

// GetMtopAppkey MtopAppkey Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetMtopAppkey() string {
	return r._mtopAppkey
}

// SetUserType is UserType Setter
// 注册用户类型
func (r *AlibabaalisportspassportaccounttokenvalidateAPIRequest) SetUserType(_userType int64) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaalisportspassportaccounttokenvalidateAPIRequest) GetUserType() int64 {
	return r._userType
}

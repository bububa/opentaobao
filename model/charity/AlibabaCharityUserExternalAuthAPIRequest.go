package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthAPIRequest 外部用户授权 API请求
// alibaba.charity.user.external.auth
//
// 外部用户授权
type AlibabaCharityUserExternalAuthAPIRequest struct {
	model.Params
	// 头像
	_avatar string
	// 昵称
	_nick string
	// 用户ID
	_userKey string
	// 用户类型
	_userType string
	// 授权范围
	_scopeId int64
}

// NewAlibabaCharityUserExternalAuthRequest 初始化AlibabaCharityUserExternalAuthAPIRequest对象
func NewAlibabaCharityUserExternalAuthRequest() *AlibabaCharityUserExternalAuthAPIRequest {
	return &AlibabaCharityUserExternalAuthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityUserExternalAuthAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.user.external.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityUserExternalAuthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityUserExternalAuthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAvatar is Avatar Setter
// 头像
func (r *AlibabaCharityUserExternalAuthAPIRequest) SetAvatar(_avatar string) error {
	r._avatar = _avatar
	r.Set("avatar", _avatar)
	return nil
}

// GetAvatar Avatar Getter
func (r AlibabaCharityUserExternalAuthAPIRequest) GetAvatar() string {
	return r._avatar
}

// SetNick is Nick Setter
// 昵称
func (r *AlibabaCharityUserExternalAuthAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r AlibabaCharityUserExternalAuthAPIRequest) GetNick() string {
	return r._nick
}

// SetUserKey is UserKey Setter
// 用户ID
func (r *AlibabaCharityUserExternalAuthAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabaCharityUserExternalAuthAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabaCharityUserExternalAuthAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaCharityUserExternalAuthAPIRequest) GetUserType() string {
	return r._userType
}

// SetScopeId is ScopeId Setter
// 授权范围
func (r *AlibabaCharityUserExternalAuthAPIRequest) SetScopeId(_scopeId int64) error {
	r._scopeId = _scopeId
	r.Set("scope_id", _scopeId)
	return nil
}

// GetScopeId ScopeId Getter
func (r AlibabaCharityUserExternalAuthAPIRequest) GetScopeId() int64 {
	return r._scopeId
}

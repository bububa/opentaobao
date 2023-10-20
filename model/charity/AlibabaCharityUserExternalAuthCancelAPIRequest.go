package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthCancelAPIRequest 取消外部授权 API请求
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
type AlibabaCharityUserExternalAuthCancelAPIRequest struct {
	model.Params
	// 用户ID
	_userKey string
	// 用户类型
	_userType string
	// 授权范围
	_scopeId int64
}

// NewAlibabaCharityUserExternalAuthCancelRequest 初始化AlibabaCharityUserExternalAuthCancelAPIRequest对象
func NewAlibabaCharityUserExternalAuthCancelRequest() *AlibabaCharityUserExternalAuthCancelAPIRequest {
	return &AlibabaCharityUserExternalAuthCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.user.external.auth.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserKey is UserKey Setter
// 用户ID
func (r *AlibabaCharityUserExternalAuthCancelAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabaCharityUserExternalAuthCancelAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetUserType() string {
	return r._userType
}

// SetScopeId is ScopeId Setter
// 授权范围
func (r *AlibabaCharityUserExternalAuthCancelAPIRequest) SetScopeId(_scopeId int64) error {
	r._scopeId = _scopeId
	r.Set("scope_id", _scopeId)
	return nil
}

// GetScopeId ScopeId Getter
func (r AlibabaCharityUserExternalAuthCancelAPIRequest) GetScopeId() int64 {
	return r._scopeId
}

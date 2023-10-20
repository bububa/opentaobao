package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharityuserexternalauthcancelAPIRequest 取消外部授权 API请求
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
type AlibabacharityuserexternalauthcancelAPIRequest struct {
	model.Params
	// 用户ID
	_userKey string
	// 用户类型
	_userType string
	// 授权范围
	_scopeId int64
}

// NewAlibabacharityuserexternalauthcancelRequest 初始化AlibabacharityuserexternalauthcancelAPIRequest对象
func NewAlibabacharityuserexternalauthcancelRequest() *AlibabacharityuserexternalauthcancelAPIRequest {
	return &AlibabacharityuserexternalauthcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.user.external.auth.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserKey is UserKey Setter
// 用户ID
func (r *AlibabacharityuserexternalauthcancelAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabacharityuserexternalauthcancelAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetUserType() string {
	return r._userType
}

// SetScopeId is ScopeId Setter
// 授权范围
func (r *AlibabacharityuserexternalauthcancelAPIRequest) SetScopeId(_scopeId int64) error {
	r._scopeId = _scopeId
	r.Set("scope_id", _scopeId)
	return nil
}

// GetScopeId ScopeId Getter
func (r AlibabacharityuserexternalauthcancelAPIRequest) GetScopeId() int64 {
	return r._scopeId
}

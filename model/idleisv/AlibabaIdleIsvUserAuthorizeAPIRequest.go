package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserAuthorizeAPIRequest 用户授权接口 API请求
// alibaba.idle.isv.user.authorize
//
// 用户授权接口
// 接口调用相关参考文档
// https://www.yuque.com/docs/share/9cd991b7-e3a3-40b6-948c-1835422d0164?# 《闲鱼优品2.0API接入说明》
type AlibabaIdleIsvUserAuthorizeAPIRequest struct {
	model.Params
	// 添加用户授权信息
	_addUserAuthorizationCmd *AddUserAuthorizationCmd
}

// NewAlibabaIdleIsvUserAuthorizeRequest 初始化AlibabaIdleIsvUserAuthorizeAPIRequest对象
func NewAlibabaIdleIsvUserAuthorizeRequest() *AlibabaIdleIsvUserAuthorizeAPIRequest {
	return &AlibabaIdleIsvUserAuthorizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvUserAuthorizeAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.user.authorize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvUserAuthorizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvUserAuthorizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddUserAuthorizationCmd is AddUserAuthorizationCmd Setter
// 添加用户授权信息
func (r *AlibabaIdleIsvUserAuthorizeAPIRequest) SetAddUserAuthorizationCmd(_addUserAuthorizationCmd *AddUserAuthorizationCmd) error {
	r._addUserAuthorizationCmd = _addUserAuthorizationCmd
	r.Set("add_user_authorization_cmd", _addUserAuthorizationCmd)
	return nil
}

// GetAddUserAuthorizationCmd AddUserAuthorizationCmd Getter
func (r AlibabaIdleIsvUserAuthorizeAPIRequest) GetAddUserAuthorizationCmd() *AddUserAuthorizationCmd {
	return r._addUserAuthorizationCmd
}

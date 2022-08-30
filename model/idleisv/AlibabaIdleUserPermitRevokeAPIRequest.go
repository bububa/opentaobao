package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitRevokeAPIRequest 删除服务商与卖家之间的订单消息绑定关系 API请求
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
type AlibabaIdleUserPermitRevokeAPIRequest struct {
	model.Params
	// 撤销用户授权请求参数
	_userPermitCmd *UserPermitCmd
}

// NewAlibabaIdleUserPermitRevokeRequest 初始化AlibabaIdleUserPermitRevokeAPIRequest对象
func NewAlibabaIdleUserPermitRevokeRequest() *AlibabaIdleUserPermitRevokeAPIRequest {
	return &AlibabaIdleUserPermitRevokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserPermitCmd is UserPermitCmd Setter
// 撤销用户授权请求参数
func (r *AlibabaIdleUserPermitRevokeAPIRequest) SetUserPermitCmd(_userPermitCmd *UserPermitCmd) error {
	r._userPermitCmd = _userPermitCmd
	r.Set("user_permit_cmd", _userPermitCmd)
	return nil
}

// GetUserPermitCmd UserPermitCmd Getter
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetUserPermitCmd() *UserPermitCmd {
	return r._userPermitCmd
}

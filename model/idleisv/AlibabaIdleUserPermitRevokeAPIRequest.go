package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleuserpermitrevokeAPIRequest 删除服务商与卖家之间的订单消息绑定关系 API请求
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
type AlibabaidleuserpermitrevokeAPIRequest struct {
	model.Params
	// 撤销用户授权请求参数
	_userPermitCmd *UserPermitCmd
}

// NewAlibabaidleuserpermitrevokeRequest 初始化AlibabaidleuserpermitrevokeAPIRequest对象
func NewAlibabaidleuserpermitrevokeRequest() *AlibabaidleuserpermitrevokeAPIRequest {
	return &AlibabaidleuserpermitrevokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleuserpermitrevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleuserpermitrevokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleuserpermitrevokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserPermitCmd is UserPermitCmd Setter
// 撤销用户授权请求参数
func (r *AlibabaidleuserpermitrevokeAPIRequest) SetUserPermitCmd(_userPermitCmd *UserPermitCmd) error {
	r._userPermitCmd = _userPermitCmd
	r.Set("user_permit_cmd", _userPermitCmd)
	return nil
}

// GetUserPermitCmd UserPermitCmd Getter
func (r AlibabaidleuserpermitrevokeAPIRequest) GetUserPermitCmd() *UserPermitCmd {
	return r._userPermitCmd
}

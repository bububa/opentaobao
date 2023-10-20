package idleisv

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleUserPermitRevokeAPIRequest) Reset() {
	r._userPermitCmd = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleUserPermitRevokeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleUserPermitRevokeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleUserPermitRevokeRequest()
	},
}

// GetAlibabaIdleUserPermitRevokeRequest 从 sync.Pool 获取 AlibabaIdleUserPermitRevokeAPIRequest
func GetAlibabaIdleUserPermitRevokeAPIRequest() *AlibabaIdleUserPermitRevokeAPIRequest {
	return poolAlibabaIdleUserPermitRevokeAPIRequest.Get().(*AlibabaIdleUserPermitRevokeAPIRequest)
}

// ReleaseAlibabaIdleUserPermitRevokeAPIRequest 将 AlibabaIdleUserPermitRevokeAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleUserPermitRevokeAPIRequest(v *AlibabaIdleUserPermitRevokeAPIRequest) {
	v.Reset()
	poolAlibabaIdleUserPermitRevokeAPIRequest.Put(v)
}

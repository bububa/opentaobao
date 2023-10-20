package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclUserpermissionsRevokeAPIRequest 回收账户权限 API请求
// alibaba.mozi.acl.userpermissions.revoke
//
// 调用此接口，会根据入参回收该账户下的该批权限点
type AlibabaMoziAclUserpermissionsRevokeAPIRequest struct {
	model.Params
	// 回收权限入参对象
	_revokePermission *RevokePermissionsRequest
}

// NewAlibabaMoziAclUserpermissionsRevokeRequest 初始化AlibabaMoziAclUserpermissionsRevokeAPIRequest对象
func NewAlibabaMoziAclUserpermissionsRevokeRequest() *AlibabaMoziAclUserpermissionsRevokeAPIRequest {
	return &AlibabaMoziAclUserpermissionsRevokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclUserpermissionsRevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.userpermissions.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclUserpermissionsRevokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclUserpermissionsRevokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRevokePermission is RevokePermission Setter
// 回收权限入参对象
func (r *AlibabaMoziAclUserpermissionsRevokeAPIRequest) SetRevokePermission(_revokePermission *RevokePermissionsRequest) error {
	r._revokePermission = _revokePermission
	r.Set("revoke_permission", _revokePermission)
	return nil
}

// GetRevokePermission RevokePermission Getter
func (r AlibabaMoziAclUserpermissionsRevokeAPIRequest) GetRevokePermission() *RevokePermissionsRequest {
	return r._revokePermission
}

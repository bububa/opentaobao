package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclUserrolesRevokeAPIRequest 回收账户被授予的角色接口 API请求
// alibaba.mozi.acl.userroles.revoke
//
// 调用此接口，会根据入参回收该账户下的该批角色
type AlibabaMoziAclUserrolesRevokeAPIRequest struct {
	model.Params
	// 回收角色入参对象
	_revokeRolesRequest *RevokeRolesRequest
}

// NewAlibabaMoziAclUserrolesRevokeRequest 初始化AlibabaMoziAclUserrolesRevokeAPIRequest对象
func NewAlibabaMoziAclUserrolesRevokeRequest() *AlibabaMoziAclUserrolesRevokeAPIRequest {
	return &AlibabaMoziAclUserrolesRevokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclUserrolesRevokeAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.userroles.revoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclUserrolesRevokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclUserrolesRevokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRevokeRolesRequest is RevokeRolesRequest Setter
// 回收角色入参对象
func (r *AlibabaMoziAclUserrolesRevokeAPIRequest) SetRevokeRolesRequest(_revokeRolesRequest *RevokeRolesRequest) error {
	r._revokeRolesRequest = _revokeRolesRequest
	r.Set("revoke_roles_request", _revokeRolesRequest)
	return nil
}

// GetRevokeRolesRequest RevokeRolesRequest Getter
func (r AlibabaMoziAclUserrolesRevokeAPIRequest) GetRevokeRolesRequest() *RevokeRolesRequest {
	return r._revokeRolesRequest
}

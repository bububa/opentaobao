package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclRoleRemoveAPIRequest
删除角色 API请求
alibaba.mozi.acl.role.remove

根据传入的角色code、租户id，删除租户内对应的角色 */
type AlibabaMoziAclRoleRemoveAPIRequest struct {
	model.Params
	// 删除角色请求对象
	_deleteRolesRequest *DeleteRolesRequest
}

// NewAlibabaMoziAclRoleRemoveRequest 初始化AlibabaMoziAclRoleRemoveAPIRequest对象
func NewAlibabaMoziAclRoleRemoveRequest() *AlibabaMoziAclRoleRemoveAPIRequest {
	return &AlibabaMoziAclRoleRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeleteRolesRequest Setter
// 删除角色请求对象
func (r *AlibabaMoziAclRoleRemoveAPIRequest) SetDeleteRolesRequest(_deleteRolesRequest *DeleteRolesRequest) error {
	r._deleteRolesRequest = _deleteRolesRequest
	r.Set("delete_roles_request", _deleteRolesRequest)
	return nil
}

// Get DeleteRolesRequest Getter
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetDeleteRolesRequest() *DeleteRolesRequest {
	return r._deleteRolesRequest
}

package moziacl

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleRemoveAPIRequest 删除角色 API请求
// alibaba.mozi.acl.role.remove
//
// 根据传入的角色code、租户id，删除租户内对应的角色
type AlibabaMoziAclRoleRemoveAPIRequest struct {
	model.Params
	// 删除角色请求对象
	_deleteRolesRequest *DeleteRolesRequest
}

// NewAlibabaMoziAclRoleRemoveRequest 初始化AlibabaMoziAclRoleRemoveAPIRequest对象
func NewAlibabaMoziAclRoleRemoveRequest() *AlibabaMoziAclRoleRemoveAPIRequest {
	return &AlibabaMoziAclRoleRemoveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziAclRoleRemoveAPIRequest) Reset() {
	r._deleteRolesRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteRolesRequest is DeleteRolesRequest Setter
// 删除角色请求对象
func (r *AlibabaMoziAclRoleRemoveAPIRequest) SetDeleteRolesRequest(_deleteRolesRequest *DeleteRolesRequest) error {
	r._deleteRolesRequest = _deleteRolesRequest
	r.Set("delete_roles_request", _deleteRolesRequest)
	return nil
}

// GetDeleteRolesRequest DeleteRolesRequest Getter
func (r AlibabaMoziAclRoleRemoveAPIRequest) GetDeleteRolesRequest() *DeleteRolesRequest {
	return r._deleteRolesRequest
}

var poolAlibabaMoziAclRoleRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziAclRoleRemoveRequest()
	},
}

// GetAlibabaMoziAclRoleRemoveRequest 从 sync.Pool 获取 AlibabaMoziAclRoleRemoveAPIRequest
func GetAlibabaMoziAclRoleRemoveAPIRequest() *AlibabaMoziAclRoleRemoveAPIRequest {
	return poolAlibabaMoziAclRoleRemoveAPIRequest.Get().(*AlibabaMoziAclRoleRemoveAPIRequest)
}

// ReleaseAlibabaMoziAclRoleRemoveAPIRequest 将 AlibabaMoziAclRoleRemoveAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziAclRoleRemoveAPIRequest(v *AlibabaMoziAclRoleRemoveAPIRequest) {
	v.Reset()
	poolAlibabaMoziAclRoleRemoveAPIRequest.Put(v)
}

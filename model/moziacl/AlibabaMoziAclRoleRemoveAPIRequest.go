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

// New

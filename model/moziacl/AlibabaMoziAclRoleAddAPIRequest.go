package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclRoleAddAPIRequest
新增一个角色 API请求
alibaba.mozi.acl.role.add

新增一个角色 */
type AlibabaMoziAclRoleAddAPIRequest struct {
	model.Params
	// 创建角色请求对象
	_createRoleRequest *CreateRoleRequest
}

// New

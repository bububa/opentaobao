package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclGrantGrantroleAPIRequest
将一个角色授予一个账号 API请求
alibaba.mozi.acl.grant.grantrole

根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色 */
type AlibabaMoziAclGrantGrantroleAPIRequest struct {
	model.Params
	// 整体入参对象
	_grantRolesRequest *GrantRolesRequest
}

// New

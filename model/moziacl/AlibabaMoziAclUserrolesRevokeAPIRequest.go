package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclUserrolesRevokeAPIRequest
回收账户被授予的角色接口 API请求
alibaba.mozi.acl.userroles.revoke

调用此接口，会根据入参回收该账户下的该批角色 */
type AlibabaMoziAclUserrolesRevokeAPIRequest struct {
	model.Params
	// 回收角色入参对象
	_revokeRolesRequest *RevokeRolesRequest
}

// New

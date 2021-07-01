package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclUserpermissionsRevokeAPIRequest
回收账户权限 API请求
alibaba.mozi.acl.userpermissions.revoke

调用此接口，会根据入参回收该账户下的该批权限点 */
type AlibabaMoziAclUserpermissionsRevokeAPIRequest struct {
	model.Params
	// 回收权限入参对象
	_revokePermission *RevokePermissionsRequest
}

// New

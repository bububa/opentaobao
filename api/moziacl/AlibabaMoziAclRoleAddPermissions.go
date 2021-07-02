package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclRoleAddPermissions 角色添加功能权限
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
func AlibabaMoziAclRoleAddPermissions(clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleAddPermissionsAPIRequest, session string) (*moziacl.AlibabaMoziAclRoleAddPermissionsAPIResponse, error) {
	var resp moziacl.AlibabaMoziAclRoleAddPermissionsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

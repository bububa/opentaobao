package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclRoleRemove 删除角色
// alibaba.mozi.acl.role.remove
//
// 根据传入的角色code、租户id，删除租户内对应的角色
func AlibabaMoziAclRoleRemove(clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleRemoveAPIRequest, resp *moziacl.AlibabaMoziAclRoleRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

/* AlibabaMoziAclRoleRemove
删除角色
alibaba.mozi.acl.role.remove

根据传入的角色code、租户id，删除租户内对应的角色 */
func AlibabaMoziAclRoleRemove(clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleRemoveAPIRequest, session string) (*moziacl.AlibabaMoziAclRoleRemoveAPIResponse, error) {
	var resp moziacl.AlibabaMoziAclRoleRemoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

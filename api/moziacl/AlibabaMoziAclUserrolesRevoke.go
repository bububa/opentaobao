package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclUserrolesRevoke 回收账户被授予的角色接口
// alibaba.mozi.acl.userroles.revoke
//
// 调用此接口，会根据入参回收该账户下的该批角色
func AlibabaMoziAclUserrolesRevoke(clt *core.SDKClient, req *moziacl.AlibabaMoziAclUserrolesRevokeAPIRequest, resp *moziacl.AlibabaMoziAclUserrolesRevokeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

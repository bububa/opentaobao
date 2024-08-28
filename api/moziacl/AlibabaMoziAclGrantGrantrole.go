package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclGrantGrantrole 将一个角色授予一个账号
// alibaba.mozi.acl.grant.grantrole
//
// 根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色
func AlibabaMoziAclGrantGrantrole(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclGrantGrantroleAPIRequest, resp *moziacl.AlibabaMoziAclGrantGrantroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

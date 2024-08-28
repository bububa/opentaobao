package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclUserpermissionsRevoke 回收账户权限
// alibaba.mozi.acl.userpermissions.revoke
//
// 调用此接口，会根据入参回收该账户下的该批权限点
func AlibabaMoziAclUserpermissionsRevoke(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclUserpermissionsRevokeAPIRequest, resp *moziacl.AlibabaMoziAclUserpermissionsRevokeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

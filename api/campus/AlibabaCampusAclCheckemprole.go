package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCheckemprole 校验用户是否有该角色
// alibaba.campus.acl.checkemprole
//
// 校验用户是否有该权限
func AlibabaCampusAclCheckemprole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclCheckemproleAPIRequest, resp *campus.AlibabaCampusAclCheckemproleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

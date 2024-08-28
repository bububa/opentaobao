package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGrantpermiitemtorole 权限赋予角色
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
func AlibabaCampusAclGrantpermiitemtorole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclGrantpermiitemtoroleAPIRequest, resp *campus.AlibabaCampusAclGrantpermiitemtoroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

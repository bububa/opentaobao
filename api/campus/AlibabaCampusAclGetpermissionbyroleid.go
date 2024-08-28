package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGetpermissionbyroleid 根据角色Id查询权限
// alibaba.campus.acl.getpermissionbyroleid
//
// 根据角色查询权限
func AlibabaCampusAclGetpermissionbyroleid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclGetpermissionbyroleidAPIRequest, resp *campus.AlibabaCampusAclGetpermissionbyroleidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

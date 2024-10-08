package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCancelpermiitemfromrole 取消角色和权限之间的关系
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
func AlibabaCampusAclCancelpermiitemfromrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclCancelpermiitemfromroleAPIRequest, resp *campus.AlibabaCampusAclCancelpermiitemfromroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

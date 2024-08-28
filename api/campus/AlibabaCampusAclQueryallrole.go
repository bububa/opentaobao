package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclQueryallrole 查询全部角色
// alibaba.campus.acl.queryallrole
//
// 查询全部园区
func AlibabaCampusAclQueryallrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclQueryallroleAPIRequest, resp *campus.AlibabaCampusAclQueryallroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

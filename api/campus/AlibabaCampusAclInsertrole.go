package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclInsertrole 新增角色
// alibaba.campus.acl.insertrole
//
// 新增角色
func AlibabaCampusAclInsertrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclInsertroleAPIRequest, resp *campus.AlibabaCampusAclInsertroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

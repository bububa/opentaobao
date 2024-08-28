package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewDeleteuserrole 删除管理员
// alibaba.campus.acl.new.deleteuserrole
//
// 删除管理员
func AlibabaCampusAclNewDeleteuserrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewDeleteuserroleAPIRequest, resp *campus.AlibabaCampusAclNewDeleteuserroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCancelrolesfromuser 撤销用户授予的角色
// alibaba.campus.acl.cancelrolesfromuser
//
// 撤销用户授予的角色
func AlibabaCampusAclCancelrolesfromuser(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclCancelrolesfromuserAPIRequest, resp *campus.AlibabaCampusAclCancelrolesfromuserAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

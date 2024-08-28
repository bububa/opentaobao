package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewRemoverole 删除角色
// alibaba.campus.acl.new.removerole
//
// 删除角色
func AlibabaCampusAclNewRemoverole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewRemoveroleAPIRequest, resp *campus.AlibabaCampusAclNewRemoveroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

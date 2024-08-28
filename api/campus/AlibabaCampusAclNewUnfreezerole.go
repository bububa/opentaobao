package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewUnfreezerole 解冻角色
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
func AlibabaCampusAclNewUnfreezerole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewUnfreezeroleAPIRequest, resp *campus.AlibabaCampusAclNewUnfreezeroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

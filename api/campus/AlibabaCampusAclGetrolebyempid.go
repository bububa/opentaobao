package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGetrolebyempid 根据用户查询角色
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
func AlibabaCampusAclGetrolebyempid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclGetrolebyempidAPIRequest, resp *campus.AlibabaCampusAclGetrolebyempidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

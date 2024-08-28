package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclAppGetpermisspkgs 分页获取应用的权限套餐
// alibaba.mozi.acl.app.getpermisspkgs
//
// 分页查询应用下的权限套餐列表
func AlibabaMoziAclAppGetpermisspkgs(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclAppGetpermisspkgsAPIRequest, resp *moziacl.AlibabaMoziAclAppGetpermisspkgsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewListuserbymenu 查询菜单下的人员
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
func AlibabaCampusAclNewListuserbymenu(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewListuserbymenuAPIRequest, resp *campus.AlibabaCampusAclNewListuserbymenuAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

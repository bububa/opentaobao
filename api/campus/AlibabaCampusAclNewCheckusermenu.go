package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckusermenu 校验用户是否有菜单权限
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
func AlibabaCampusAclNewCheckusermenu(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckusermenuAPIRequest, resp *campus.AlibabaCampusAclNewCheckusermenuAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

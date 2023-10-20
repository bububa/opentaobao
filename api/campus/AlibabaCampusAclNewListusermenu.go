package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewListusermenu 查询用户有权限的菜单树
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
func AlibabaCampusAclNewListusermenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewListusermenuAPIRequest, resp *campus.AlibabaCampusAclNewListusermenuAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

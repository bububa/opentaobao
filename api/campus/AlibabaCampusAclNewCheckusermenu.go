package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckusermenu 校验用户是否有菜单权限
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
func AlibabaCampusAclNewCheckusermenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckusermenuAPIRequest, resp *campus.AlibabaCampusAclNewCheckusermenuAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

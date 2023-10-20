package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCheckemprole 校验用户是否有该角色
// alibaba.campus.acl.checkemprole
//
// 校验用户是否有该权限
func AlibabaCampusAclCheckemprole(clt *core.SDKClient, req *campus.AlibabaCampusAclCheckemproleAPIRequest, resp *campus.AlibabaCampusAclCheckemproleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

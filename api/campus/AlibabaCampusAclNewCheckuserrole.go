package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckuserrole 校验用户是否有角色
// alibaba.campus.acl.new.checkuserrole
//
// 校验用户是否有角色
func AlibabaCampusAclNewCheckuserrole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserroleAPIRequest, resp *campus.AlibabaCampusAclNewCheckuserroleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

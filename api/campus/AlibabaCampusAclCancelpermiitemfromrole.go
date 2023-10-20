package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCancelpermiitemfromrole 取消角色和权限之间的关系
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
func AlibabaCampusAclCancelpermiitemfromrole(clt *core.SDKClient, req *campus.AlibabaCampusAclCancelpermiitemfromroleAPIRequest, resp *campus.AlibabaCampusAclCancelpermiitemfromroleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

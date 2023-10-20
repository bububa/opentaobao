package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclQueryallrole 查询全部角色
// alibaba.campus.acl.queryallrole
//
// 查询全部园区
func AlibabaCampusAclQueryallrole(clt *core.SDKClient, req *campus.AlibabaCampusAclQueryallroleAPIRequest, resp *campus.AlibabaCampusAclQueryallroleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

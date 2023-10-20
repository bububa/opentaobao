package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclInsertrole 新增角色
// alibaba.campus.acl.insertrole
//
// 新增角色
func AlibabaCampusAclInsertrole(clt *core.SDKClient, req *campus.AlibabaCampusAclInsertroleAPIRequest, resp *campus.AlibabaCampusAclInsertroleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

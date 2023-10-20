package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewSaverolewithmenu 保存角色级联保存角色和权限的关系
// alibaba.campus.acl.new.saverolewithmenu
//
// 保存角色级联保存角色和权限的关系
func AlibabaCampusAclNewSaverolewithmenu(clt *core.SDKClient, req *campus.AlibabaCampusAclNewSaverolewithmenuAPIRequest, resp *campus.AlibabaCampusAclNewSaverolewithmenuAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

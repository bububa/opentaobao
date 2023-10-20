package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewListuserroles 查询并标记用户选择的角色
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
func AlibabaCampusAclNewListuserroles(clt *core.SDKClient, req *campus.AlibabaCampusAclNewListuserrolesAPIRequest, resp *campus.AlibabaCampusAclNewListuserrolesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

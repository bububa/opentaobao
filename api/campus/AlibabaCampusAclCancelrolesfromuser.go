package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclCancelrolesfromuser 撤销用户授予的角色
// alibaba.campus.acl.cancelrolesfromuser
//
// 撤销用户授予的角色
func AlibabaCampusAclCancelrolesfromuser(clt *core.SDKClient, req *campus.AlibabaCampusAclCancelrolesfromuserAPIRequest, resp *campus.AlibabaCampusAclCancelrolesfromuserAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

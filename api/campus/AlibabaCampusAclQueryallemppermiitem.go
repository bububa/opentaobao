package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclQueryallemppermiitem 查询员工全部权限(包括角色下面的权限)
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
func AlibabaCampusAclQueryallemppermiitem(clt *core.SDKClient, req *campus.AlibabaCampusAclQueryallemppermiitemAPIRequest, resp *campus.AlibabaCampusAclQueryallemppermiitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

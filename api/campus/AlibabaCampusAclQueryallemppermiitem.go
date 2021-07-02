package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclQueryallemppermiitem 查询员工全部权限(包括角色下面的权限)
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
func AlibabaCampusAclQueryallemppermiitem(clt *core.SDKClient, req *campus.AlibabaCampusAclQueryallemppermiitemAPIRequest, session string) (*campus.AlibabaCampusAclQueryallemppermiitemAPIResponse, error) {
	var resp campus.AlibabaCampusAclQueryallemppermiitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

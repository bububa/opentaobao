package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclqueryallemppermiitem 查询员工全部权限(包括角色下面的权限)
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
func Alibabacampusaclqueryallemppermiitem(clt *core.SDKClient, req *campus.AlibabacampusaclqueryallemppermiitemAPIRequest, session string) (*campus.AlibabacampusaclqueryallemppermiitemAPIResponse, error) {
	var resp campus.AlibabacampusaclqueryallemppermiitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclqueryallrole 查询全部角色
// alibaba.campus.acl.queryallrole
//
// 查询全部园区
func Alibabacampusaclqueryallrole(clt *core.SDKClient, req *campus.AlibabacampusaclqueryallroleAPIRequest, session string) (*campus.AlibabacampusaclqueryallroleAPIResponse, error) {
	var resp campus.AlibabacampusaclqueryallroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

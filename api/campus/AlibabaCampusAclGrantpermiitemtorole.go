package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclgrantpermiitemtorole 权限赋予角色
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
func Alibabacampusaclgrantpermiitemtorole(clt *core.SDKClient, req *campus.AlibabacampusaclgrantpermiitemtoroleAPIRequest, session string) (*campus.AlibabacampusaclgrantpermiitemtoroleAPIResponse, error) {
	var resp campus.AlibabacampusaclgrantpermiitemtoroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

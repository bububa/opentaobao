package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewlistusermenu 查询用户有权限的菜单树
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
func Alibabacampusaclnewlistusermenu(clt *core.SDKClient, req *campus.AlibabacampusaclnewlistusermenuAPIRequest, session string) (*campus.AlibabacampusaclnewlistusermenuAPIResponse, error) {
	var resp campus.AlibabacampusaclnewlistusermenuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

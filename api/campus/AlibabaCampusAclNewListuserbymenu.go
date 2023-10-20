package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewlistuserbymenu 查询菜单下的人员
// alibaba.campus.acl.new.listuserbymenu
//
// 查询拥有菜单权限的用户
func Alibabacampusaclnewlistuserbymenu(clt *core.SDKClient, req *campus.AlibabacampusaclnewlistuserbymenuAPIRequest, session string) (*campus.AlibabacampusaclnewlistuserbymenuAPIResponse, error) {
	var resp campus.AlibabacampusaclnewlistuserbymenuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

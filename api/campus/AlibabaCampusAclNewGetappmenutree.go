package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewgetappmenutree 查询应用下的菜单树
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
func Alibabacampusaclnewgetappmenutree(clt *core.SDKClient, req *campus.AlibabacampusaclnewgetappmenutreeAPIRequest, session string) (*campus.AlibabacampusaclnewgetappmenutreeAPIResponse, error) {
	var resp campus.AlibabacampusaclnewgetappmenutreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

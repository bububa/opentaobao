package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewgetrolewithmenutreenodes 根据角色id查询权限
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
func Alibabacampusaclnewgetrolewithmenutreenodes(clt *core.SDKClient, req *campus.AlibabacampusaclnewgetrolewithmenutreenodesAPIRequest, session string) (*campus.AlibabacampusaclnewgetrolewithmenutreenodesAPIResponse, error) {
	var resp campus.AlibabacampusaclnewgetrolewithmenutreenodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

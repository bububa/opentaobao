package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclgetrolebyempid 根据用户查询角色
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
func Alibabacampusaclgetrolebyempid(clt *core.SDKClient, req *campus.AlibabacampusaclgetrolebyempidAPIRequest, session string) (*campus.AlibabacampusaclgetrolebyempidAPIResponse, error) {
	var resp campus.AlibabacampusaclgetrolebyempidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

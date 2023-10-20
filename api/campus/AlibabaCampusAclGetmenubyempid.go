package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclgetmenubyempid 查询用户的菜单
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
func Alibabacampusaclgetmenubyempid(clt *core.SDKClient, req *campus.AlibabacampusaclgetmenubyempidAPIRequest, session string) (*campus.AlibabacampusaclgetmenubyempidAPIResponse, error) {
	var resp campus.AlibabacampusaclgetmenubyempidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

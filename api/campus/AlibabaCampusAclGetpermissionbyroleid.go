package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclgetpermissionbyroleid 根据角色Id查询权限
// alibaba.campus.acl.getpermissionbyroleid
//
// 根据角色查询权限
func Alibabacampusaclgetpermissionbyroleid(clt *core.SDKClient, req *campus.AlibabacampusaclgetpermissionbyroleidAPIRequest, session string) (*campus.AlibabacampusaclgetpermissionbyroleidAPIResponse, error) {
	var resp campus.AlibabacampusaclgetpermissionbyroleidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

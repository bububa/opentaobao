package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclcancelpermiitemfromrole 取消角色和权限之间的关系
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
func Alibabacampusaclcancelpermiitemfromrole(clt *core.SDKClient, req *campus.AlibabacampusaclcancelpermiitemfromroleAPIRequest, session string) (*campus.AlibabacampusaclcancelpermiitemfromroleAPIResponse, error) {
	var resp campus.AlibabacampusaclcancelpermiitemfromroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

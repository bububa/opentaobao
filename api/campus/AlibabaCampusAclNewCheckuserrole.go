package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewcheckuserrole 校验用户是否有角色
// alibaba.campus.acl.new.checkuserrole
//
// 校验用户是否有角色
func Alibabacampusaclnewcheckuserrole(clt *core.SDKClient, req *campus.AlibabacampusaclnewcheckuserroleAPIRequest, session string) (*campus.AlibabacampusaclnewcheckuserroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewcheckuserroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

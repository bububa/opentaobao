package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewcheckuserpermission 校验用户是否有权限
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
func Alibabacampusaclnewcheckuserpermission(clt *core.SDKClient, req *campus.AlibabacampusaclnewcheckuserpermissionAPIRequest, session string) (*campus.AlibabacampusaclnewcheckuserpermissionAPIResponse, error) {
	var resp campus.AlibabacampusaclnewcheckuserpermissionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

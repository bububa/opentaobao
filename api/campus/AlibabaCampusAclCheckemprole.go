package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclcheckemprole 校验用户是否有该角色
// alibaba.campus.acl.checkemprole
//
// 校验用户是否有该权限
func Alibabacampusaclcheckemprole(clt *core.SDKClient, req *campus.AlibabacampusaclcheckemproleAPIRequest, session string) (*campus.AlibabacampusaclcheckemproleAPIResponse, error) {
	var resp campus.AlibabacampusaclcheckemproleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

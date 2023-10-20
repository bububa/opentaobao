package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewcheckusermenu 校验用户是否有菜单权限
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
func Alibabacampusaclnewcheckusermenu(clt *core.SDKClient, req *campus.AlibabacampusaclnewcheckusermenuAPIRequest, session string) (*campus.AlibabacampusaclnewcheckusermenuAPIResponse, error) {
	var resp campus.AlibabacampusaclnewcheckusermenuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

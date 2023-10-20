package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewlistuserroles 查询并标记用户选择的角色
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
func Alibabacampusaclnewlistuserroles(clt *core.SDKClient, req *campus.AlibabacampusaclnewlistuserrolesAPIRequest, session string) (*campus.AlibabacampusaclnewlistuserrolesAPIResponse, error) {
	var resp campus.AlibabacampusaclnewlistuserrolesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

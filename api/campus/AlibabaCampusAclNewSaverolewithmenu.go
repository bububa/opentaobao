package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewsaverolewithmenu 保存角色级联保存角色和权限的关系
// alibaba.campus.acl.new.saverolewithmenu
//
// 保存角色级联保存角色和权限的关系
func Alibabacampusaclnewsaverolewithmenu(clt *core.SDKClient, req *campus.AlibabacampusaclnewsaverolewithmenuAPIRequest, session string) (*campus.AlibabacampusaclnewsaverolewithmenuAPIResponse, error) {
	var resp campus.AlibabacampusaclnewsaverolewithmenuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

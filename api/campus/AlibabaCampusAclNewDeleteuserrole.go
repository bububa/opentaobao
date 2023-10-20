package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewdeleteuserrole 删除管理员
// alibaba.campus.acl.new.deleteuserrole
//
// 删除管理员
func Alibabacampusaclnewdeleteuserrole(clt *core.SDKClient, req *campus.AlibabacampusaclnewdeleteuserroleAPIRequest, session string) (*campus.AlibabacampusaclnewdeleteuserroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewdeleteuserroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

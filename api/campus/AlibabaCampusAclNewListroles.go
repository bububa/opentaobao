package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewlistroles 查询全部角色
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
func Alibabacampusaclnewlistroles(clt *core.SDKClient, req *campus.AlibabacampusaclnewlistrolesAPIRequest, session string) (*campus.AlibabacampusaclnewlistrolesAPIResponse, error) {
	var resp campus.AlibabacampusaclnewlistrolesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

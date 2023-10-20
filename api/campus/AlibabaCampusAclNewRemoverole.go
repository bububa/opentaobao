package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewremoverole 删除角色
// alibaba.campus.acl.new.removerole
//
// 删除角色
func Alibabacampusaclnewremoverole(clt *core.SDKClient, req *campus.AlibabacampusaclnewremoveroleAPIRequest, session string) (*campus.AlibabacampusaclnewremoveroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewremoveroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

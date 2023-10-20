package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewunfreezerole 解冻角色
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
func Alibabacampusaclnewunfreezerole(clt *core.SDKClient, req *campus.AlibabacampusaclnewunfreezeroleAPIRequest, session string) (*campus.AlibabacampusaclnewunfreezeroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewunfreezeroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

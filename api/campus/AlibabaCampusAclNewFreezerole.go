package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclnewfreezerole 冻结角色
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
func Alibabacampusaclnewfreezerole(clt *core.SDKClient, req *campus.AlibabacampusaclnewfreezeroleAPIRequest, session string) (*campus.AlibabacampusaclnewfreezeroleAPIResponse, error) {
	var resp campus.AlibabacampusaclnewfreezeroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

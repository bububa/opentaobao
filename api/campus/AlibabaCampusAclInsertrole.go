package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclinsertrole 新增角色
// alibaba.campus.acl.insertrole
//
// 新增角色
func Alibabacampusaclinsertrole(clt *core.SDKClient, req *campus.AlibabacampusaclinsertroleAPIRequest, session string) (*campus.AlibabacampusaclinsertroleAPIResponse, error) {
	var resp campus.AlibabacampusaclinsertroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

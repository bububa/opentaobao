package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclupdategrantroletouser 修改用户到角色关系
// alibaba.campus.acl.updategrantroletouser
//
// 修改用户到角色关系
func Alibabacampusaclupdategrantroletouser(clt *core.SDKClient, req *campus.AlibabacampusaclupdategrantroletouserAPIRequest, session string) (*campus.AlibabacampusaclupdategrantroletouserAPIResponse, error) {
	var resp campus.AlibabacampusaclupdategrantroletouserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

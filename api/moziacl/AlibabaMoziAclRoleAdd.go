package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclroleadd 新增一个角色
// alibaba.mozi.acl.role.add
//
// 新增一个角色
func Alibabamoziaclroleadd(clt *core.SDKClient, req *moziacl.AlibabamoziaclroleaddAPIRequest, session string) (*moziacl.AlibabamoziaclroleaddAPIResponse, error) {
	var resp moziacl.AlibabamoziaclroleaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

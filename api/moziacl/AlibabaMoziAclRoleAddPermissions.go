package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclroleaddpermissions 角色添加功能权限
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
func Alibabamoziaclroleaddpermissions(clt *core.SDKClient, req *moziacl.AlibabamoziaclroleaddpermissionsAPIRequest, session string) (*moziacl.AlibabamoziaclroleaddpermissionsAPIResponse, error) {
	var resp moziacl.AlibabamoziaclroleaddpermissionsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

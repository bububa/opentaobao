package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclroleremovepermissions 角色移除功能权限
// alibaba.mozi.acl.role.remove.permissions
//
// 从角色中移除一批功能权限
func Alibabamoziaclroleremovepermissions(clt *core.SDKClient, req *moziacl.AlibabamoziaclroleremovepermissionsAPIRequest, session string) (*moziacl.AlibabamoziaclroleremovepermissionsAPIResponse, error) {
	var resp moziacl.AlibabamoziaclroleremovepermissionsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclpermissionpkgaddpermissions 权限套餐添加权限
// alibaba.mozi.acl.permissionpkg.add.permissions
//
// 此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
func Alibabamoziaclpermissionpkgaddpermissions(clt *core.SDKClient, req *moziacl.AlibabamoziaclpermissionpkgaddpermissionsAPIRequest, session string) (*moziacl.AlibabamoziaclpermissionpkgaddpermissionsAPIResponse, error) {
	var resp moziacl.AlibabamoziaclpermissionpkgaddpermissionsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

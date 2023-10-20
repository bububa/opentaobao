package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclpermissionpkgaddroles 将角色添加到权限套餐中
// alibaba.mozi.acl.permissionpkg.add.roles
//
// 此接口是将应用下的一批角色添加到该应用的某个权限套餐中
func Alibabamoziaclpermissionpkgaddroles(clt *core.SDKClient, req *moziacl.AlibabamoziaclpermissionpkgaddrolesAPIRequest, session string) (*moziacl.AlibabamoziaclpermissionpkgaddrolesAPIResponse, error) {
	var resp moziacl.AlibabamoziaclpermissionpkgaddrolesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

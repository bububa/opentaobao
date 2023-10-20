package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// Alibabamoziaclpermissionpagerolepermission 分页查询角色下包含的权限列表
// alibaba.mozi.acl.permission.page.rolepermission
//
// 根据传入的角色name，分页查询该角色包含的权限列表
func Alibabamoziaclpermissionpagerolepermission(clt *core.SDKClient, req *moziacl.AlibabamoziaclpermissionpagerolepermissionAPIRequest, session string) (*moziacl.AlibabamoziaclpermissionpagerolepermissionAPIResponse, error) {
	var resp moziacl.AlibabamoziaclpermissionpagerolepermissionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

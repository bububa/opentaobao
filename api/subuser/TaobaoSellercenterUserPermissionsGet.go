package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// Taobaosellercenteruserpermissionsget 获取指定用户的权限集合
// taobao.sellercenter.user.permissions.get
//
// 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
func Taobaosellercenteruserpermissionsget(clt *core.SDKClient, req *subuser.TaobaosellercenteruserpermissionsgetAPIRequest, session string) (*subuser.TaobaosellercenteruserpermissionsgetAPIResponse, error) {
	var resp subuser.TaobaosellercenteruserpermissionsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

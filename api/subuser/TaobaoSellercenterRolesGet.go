package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// Taobaosellercenterrolesget 获取指定卖家的角色列表
// taobao.sellercenter.roles.get
//
// 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
func Taobaosellercenterrolesget(clt *core.SDKClient, req *subuser.TaobaosellercenterrolesgetAPIRequest, session string) (*subuser.TaobaosellercenterrolesgetAPIResponse, error) {
	var resp subuser.TaobaosellercenterrolesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

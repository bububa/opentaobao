package subuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSellercenterRolesGet 获取指定卖家的角色列表
// taobao.sellercenter.roles.get
//
// 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
func TaobaoSellercenterRolesGet(ctx context.Context, clt *core.SDKClient, req *subuser.TaobaoSellercenterRolesGetAPIRequest, resp *subuser.TaobaoSellercenterRolesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

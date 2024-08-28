package subuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersGet 获取指定账户的子账号简易信息列表
// taobao.subusers.get
//
// 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
func TaobaoSubusersGet(ctx context.Context, clt *core.SDKClient, req *subuser.TaobaoSubusersGetAPIRequest, resp *subuser.TaobaoSubusersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

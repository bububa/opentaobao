package subuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubusersPage 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口）
// taobao.subusers.page
//
// 分页获取指定账户的子账号简易信息列表
// （新isv接入建议使用taobao.sellercenter.subusers.page接口）
func TaobaoSubusersPage(ctx context.Context, clt *core.SDKClient, req *subuser.TaobaoSubusersPageAPIRequest, resp *subuser.TaobaoSubusersPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

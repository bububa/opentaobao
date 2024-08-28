package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgNewuserOrderGet 淘宝客-推广者-新用户订单明细查询
// taobao.tbk.dg.newuser.order.get
//
// 拉新API
func TaobaoTbkDgNewuserOrderGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgNewuserOrderGetAPIRequest, resp *tbk.TaobaoTbkDgNewuserOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

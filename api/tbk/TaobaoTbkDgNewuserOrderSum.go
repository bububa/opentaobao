package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgNewuserOrderSum 淘宝客-推广者-拉新活动对应数据查询
// taobao.tbk.dg.newuser.order.sum
//
// 拉新活动汇总API
func TaobaoTbkDgNewuserOrderSum(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgNewuserOrderSumAPIRequest, resp *tbk.TaobaoTbkDgNewuserOrderSumAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

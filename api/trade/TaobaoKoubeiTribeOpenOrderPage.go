package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoKoubeiTribeOpenOrderPage 口碑综合体订单列表信息查询
// taobao.koubei.tribe.open.order.page
//
// 查询口碑商圈用户的订单列表信息
func TaobaoKoubeiTribeOpenOrderPage(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoKoubeiTribeOpenOrderPageAPIRequest, resp *trade.TaobaoKoubeiTribeOpenOrderPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryAdjustTrade 交易库存调整单
// taobao.inventory.adjust.trade
//
// 商家交易调整库存，淘宝交易、B2B经销等
func TaobaoInventoryAdjustTrade(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryAdjustTradeAPIRequest, resp *fenxiao.TaobaoInventoryAdjustTradeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

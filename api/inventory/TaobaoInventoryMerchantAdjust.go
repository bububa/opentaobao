package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryMerchantAdjust 货品库存商家端调整
// taobao.inventory.merchant.adjust
//
// 货品库存商家端调整 ，入库，出库，盘点
func TaobaoInventoryMerchantAdjust(ctx context.Context, clt *core.SDKClient, req *inventory.TaobaoInventoryMerchantAdjustAPIRequest, resp *inventory.TaobaoInventoryMerchantAdjustAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

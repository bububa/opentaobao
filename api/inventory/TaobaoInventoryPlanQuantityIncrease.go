package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanQuantityIncrease 计划库存的增量编辑
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
func TaobaoInventoryPlanQuantityIncrease(ctx context.Context, clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQuantityIncreaseAPIRequest, resp *inventory.TaobaoInventoryPlanQuantityIncreaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

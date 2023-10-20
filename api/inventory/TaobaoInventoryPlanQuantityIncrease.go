package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanQuantityIncrease 计划库存的增量编辑
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
func TaobaoInventoryPlanQuantityIncrease(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQuantityIncreaseAPIRequest, resp *inventory.TaobaoInventoryPlanQuantityIncreaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

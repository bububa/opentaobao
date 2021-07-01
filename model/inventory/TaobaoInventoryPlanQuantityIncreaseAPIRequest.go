package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryPlanQuantityIncreaseAPIRequest
计划库存的增量编辑 API请求
taobao.inventory.plan.quantity.increase

计划库存的增量编辑 */
type TaobaoInventoryPlanQuantityIncreaseAPIRequest struct {
	model.Params
	// 增量编辑计划库存入参
	_planInvAdjustTop *PlanInvAdjustTopDto
}

// New

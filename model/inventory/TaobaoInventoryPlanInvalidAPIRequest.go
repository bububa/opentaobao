package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryPlanInvalidAPIRequest
失效计划库存 API请求
taobao.inventory.plan.invalid

计划库存的失效服务 */
type TaobaoInventoryPlanInvalidAPIRequest struct {
	model.Params
	// 计划库存失效入参
	_planStopTop *PlanStopTopDto
}

// New

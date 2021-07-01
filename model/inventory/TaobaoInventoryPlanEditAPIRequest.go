package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryPlanEditAPIRequest
设置计划库存 API请求
taobao.inventory.plan.edit

初始化计划库存，或者编辑已经存在的计划库存 */
type TaobaoInventoryPlanEditAPIRequest struct {
	model.Params
	// 计划库存设置入参
	_planTop *PlanTopDto
}

// New

package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryPlanQueryAPIRequest
计划库存查询 API请求
taobao.inventory.plan.query

计划库存查询 */
type TaobaoInventoryPlanQueryAPIRequest struct {
	model.Params
	// 计划库存查询入参
	_param *InvUnifyPlanTopQuerys
}

// New

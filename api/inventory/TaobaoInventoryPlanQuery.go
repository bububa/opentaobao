package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanQuery 计划库存查询
// taobao.inventory.plan.query
//
// 计划库存查询
func TaobaoInventoryPlanQuery(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQueryAPIRequest, resp *inventory.TaobaoInventoryPlanQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

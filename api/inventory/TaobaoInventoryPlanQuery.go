package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanQuery 计划库存查询
// taobao.inventory.plan.query
//
// 计划库存查询
func TaobaoInventoryPlanQuery(clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQueryAPIRequest, session string) (*inventory.TaobaoInventoryPlanQueryAPIResponse, error) {
	var resp inventory.TaobaoInventoryPlanQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

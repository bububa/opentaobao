package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanQuery 计划库存查询
// taobao.inventory.plan.query
//
// 计划库存查询
func TaobaoInventoryPlanQuery(ctx context.Context, clt *core.SDKClient, req *inventory.TaobaoInventoryPlanQueryAPIRequest, resp *inventory.TaobaoInventoryPlanQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

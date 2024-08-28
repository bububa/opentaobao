package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanInvalid 失效计划库存
// taobao.inventory.plan.invalid
//
// 计划库存的失效服务
func TaobaoInventoryPlanInvalid(ctx context.Context, clt *core.SDKClient, req *inventory.TaobaoInventoryPlanInvalidAPIRequest, resp *inventory.TaobaoInventoryPlanInvalidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

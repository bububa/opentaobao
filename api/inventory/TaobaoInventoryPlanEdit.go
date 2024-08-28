package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryPlanEdit 设置计划库存
// taobao.inventory.plan.edit
//
// 初始化计划库存，或者编辑已经存在的计划库存
func TaobaoInventoryPlanEdit(ctx context.Context, clt *core.SDKClient, req *inventory.TaobaoInventoryPlanEditAPIRequest, resp *inventory.TaobaoInventoryPlanEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCapacityRuleUpdate 仓产能创建/更新
// taobao.logistics.warehouse.capacity.rule.update
//
// 仓产能创建/更新
func TaobaoLogisticsWarehouseCapacityRuleUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIRequest, resp *ascp.TaobaoLogisticsWarehouseCapacityRuleUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

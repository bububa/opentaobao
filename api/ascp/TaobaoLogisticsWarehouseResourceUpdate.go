package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseResourceUpdate 服务商新建/更新仓
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
func TaobaoLogisticsWarehouseResourceUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseResourceUpdateAPIRequest, resp *ascp.TaobaoLogisticsWarehouseResourceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

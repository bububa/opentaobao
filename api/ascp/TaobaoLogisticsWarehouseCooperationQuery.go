package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationQuery 仓合作关系查询
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
func TaobaoLogisticsWarehouseCooperationQuery(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationQueryAPIRequest, resp *ascp.TaobaoLogisticsWarehouseCooperationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

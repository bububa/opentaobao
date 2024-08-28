package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsWarehouseCooperationBatchConfirm 仓合作关系确认
// taobao.logistics.warehouse.cooperation.batch.confirm
//
// 仓合作关系确认
func TaobaoLogisticsWarehouseCooperationBatchConfirm(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsWarehouseCooperationBatchConfirmAPIRequest, resp *ascp.TaobaoLogisticsWarehouseCooperationBatchConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

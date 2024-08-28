package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsDeliveryLineBatchDelete 线路能力删除
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
func TaobaoLogisticsDeliveryLineBatchDelete(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest, resp *ascp.TaobaoLogisticsDeliveryLineBatchDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

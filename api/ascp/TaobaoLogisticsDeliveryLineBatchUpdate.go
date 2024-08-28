package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsDeliveryLineBatchUpdate 线路能力创建/更新
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
func TaobaoLogisticsDeliveryLineBatchUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest, resp *ascp.TaobaoLogisticsDeliveryLineBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

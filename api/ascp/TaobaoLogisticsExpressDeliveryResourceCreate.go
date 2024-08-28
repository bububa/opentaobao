package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressDeliveryResourceCreate 新建/更新配资源
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
func TaobaoLogisticsExpressDeliveryResourceCreate(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest, resp *ascp.TaobaoLogisticsExpressDeliveryResourceCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

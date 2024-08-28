package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxWarehouseInboundCallback 翱象入库回传
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
func AlibabaAxWarehouseInboundCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAxWarehouseInboundCallbackAPIRequest, resp *wdk.AlibabaAxWarehouseInboundCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxWarehouseOutboundCallback 翱象出仓回传
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
func AlibabaAxWarehouseOutboundCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAxWarehouseOutboundCallbackAPIRequest, resp *wdk.AlibabaAxWarehouseOutboundCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

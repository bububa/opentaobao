package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWarehouseOrderCancel 仓作业取消下发
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
func WdkWarehouseOrderCancel(ctx context.Context, clt *core.SDKClient, req *wdk.WdkWarehouseOrderCancelAPIRequest, resp *wdk.WdkWarehouseOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

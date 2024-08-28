package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillWarehouseWorkOrderCallback 标准化仓作业单回传接口
// alibaba.wdk.fulfill.warehouse.work.order.callback
//
// 标准化仓作业单回传接口
func AlibabaWdkFulfillWarehouseWorkOrderCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

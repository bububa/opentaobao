package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallback 大润发B2C仓作业单回传接口
// alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback
//
// 大润发B2C仓作业单回传接口
func AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

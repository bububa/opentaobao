package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillDmsDeliveryWorkOrderCallback 末端配配送作业回传
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
func AlibabaWdkFulfillDmsDeliveryWorkOrderCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillDmsEbeecakeWorkOrderCallback 北京小蜜蜂配作业回传
// alibaba.wdk.fulfill.dms.ebeecake.work.order.callback
//
// 北京小蜜蜂配作业回传。
func AlibabaWdkFulfillDmsEbeecakeWorkOrderCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillDmsEbeecakeWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

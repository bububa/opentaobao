package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillBatchOnTaskStatusChanged 物流管控作业状态回传
// alibaba.wdk.fulfill.batch.on.task.status.changed
//
// 物流管控作业状态回传
func AlibabaWdkFulfillBatchOnTaskStatusChanged(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest, resp *wdk.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

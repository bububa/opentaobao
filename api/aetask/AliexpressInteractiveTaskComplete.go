package aetask

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// AliexpressInteractiveTaskComplete 任务完成接口
// aliexpress.interactive.task.complete
//
// 用户完成任务
func AliexpressInteractiveTaskComplete(ctx context.Context, clt *core.SDKClient, req *aetask.AliexpressInteractiveTaskCompleteAPIRequest, resp *aetask.AliexpressInteractiveTaskCompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

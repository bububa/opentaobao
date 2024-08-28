package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetask 领取任务
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
func AlibabaAlscGrowthInteractiveTaskReceivetask(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

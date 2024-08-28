package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskprize 任务领奖
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
func AlibabaAlscGrowthInteractiveTaskReceivetaskprize(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

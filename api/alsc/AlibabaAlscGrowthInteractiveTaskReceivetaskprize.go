package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskprize 任务领奖
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
func AlibabaAlscGrowthInteractiveTaskReceivetaskprize(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

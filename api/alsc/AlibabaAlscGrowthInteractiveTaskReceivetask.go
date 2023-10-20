package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetask 领取任务
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
func AlibabaAlscGrowthInteractiveTaskReceivetask(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

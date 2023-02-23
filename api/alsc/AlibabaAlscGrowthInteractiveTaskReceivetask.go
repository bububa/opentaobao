package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetask 领取任务
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
func AlibabaAlscGrowthInteractiveTaskReceivetask(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

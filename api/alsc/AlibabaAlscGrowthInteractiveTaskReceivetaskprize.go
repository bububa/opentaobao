package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskprize 任务领奖
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
func AlibabaAlscGrowthInteractiveTaskReceivetaskprize(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

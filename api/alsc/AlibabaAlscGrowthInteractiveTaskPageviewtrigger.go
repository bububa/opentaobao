package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskPageviewtrigger 浏览打点接口
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
func AlibabaAlscGrowthInteractiveTaskPageviewtrigger(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

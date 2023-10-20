package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskPageviewtrigger 浏览打点接口
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
func AlibabaAlscGrowthInteractiveTaskPageviewtrigger(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest, resp *alsc.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

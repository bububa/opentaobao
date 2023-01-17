package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveTaskQuerytask 本地生活用户增长互动任务平台查询任务
// alibaba.alsc.growth.interactive.task.querytask
//
// 本地生活用户增长互动任务平台查询任务
func AlibabaAlscGrowthInteractiveTaskQuerytask(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveTaskQuerytaskAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

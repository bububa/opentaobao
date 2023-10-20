package aetask

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// AliexpressInteractiveTaskComplete 任务完成接口
// aliexpress.interactive.task.complete
//
// 用户完成任务
func AliexpressInteractiveTaskComplete(clt *core.SDKClient, req *aetask.AliexpressInteractiveTaskCompleteAPIRequest, session string) (*aetask.AliexpressInteractiveTaskCompleteAPIResponse, error) {
	var resp aetask.AliexpressInteractiveTaskCompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

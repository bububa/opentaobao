package aetask

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// AliexpressInteractiveTaskComplete 任务完成接口
// aliexpress.interactive.task.complete
//
// 用户完成任务
func AliexpressInteractiveTaskComplete(clt *core.SDKClient, req *aetask.AliexpressInteractiveTaskCompleteAPIRequest, resp *aetask.AliexpressInteractiveTaskCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package aetask

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// Aliexpressinteractivetaskcomplete 任务完成接口
// aliexpress.interactive.task.complete
//
// 用户完成任务
func Aliexpressinteractivetaskcomplete(clt *core.SDKClient, req *aetask.AliexpressinteractivetaskcompleteAPIRequest, session string) (*aetask.AliexpressinteractivetaskcompleteAPIResponse, error) {
	var resp aetask.AliexpressinteractivetaskcompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

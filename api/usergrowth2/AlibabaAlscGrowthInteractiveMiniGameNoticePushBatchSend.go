package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSend 批量发送push
// alibaba.alsc.growth.interactive.mini.game.notice.push.batch.send
//
// 批量发送push
func AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSend(clt *core.SDKClient, req *usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIRequest, session string) (*usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponse, error) {
	var resp usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushBatchSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

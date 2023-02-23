package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// AlibabaAlscGrowthInteractiveMiniGameNoticePushSend 小游戏发送push
// alibaba.alsc.growth.interactive.mini.game.notice.push.send
//
// 小游戏发送push
func AlibabaAlscGrowthInteractiveMiniGameNoticePushSend(clt *core.SDKClient, req *usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIRequest, session string) (*usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIResponse, error) {
	var resp usergrowth2.AlibabaAlscGrowthInteractiveMiniGameNoticePushSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

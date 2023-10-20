package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerStop 用户停止游戏
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
func AlibabaCloudgameInteractiveGamePlayerStop(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

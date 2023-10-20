package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameHeartbeat 游戏玩家心跳
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
func AlibabaCloudgameInteractiveGameHeartbeat(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameHeartbeat 游戏玩家心跳
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
func AlibabaCloudgameInteractiveGameHeartbeat(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

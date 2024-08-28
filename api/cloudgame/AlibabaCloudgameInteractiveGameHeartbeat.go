package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameHeartbeat 游戏玩家心跳
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
func AlibabaCloudgameInteractiveGameHeartbeat(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameHeartbeatAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

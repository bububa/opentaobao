package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerKickout 踢人
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
func AlibabaCloudgameInteractiveGamePlayerKickout(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

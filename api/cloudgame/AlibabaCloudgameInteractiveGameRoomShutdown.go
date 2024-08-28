package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameRoomShutdown 强制关闭房间
// alibaba.cloudgame.interactive.game.room.shutdown
//
// 强制关闭房间
func AlibabaCloudgameInteractiveGameRoomShutdown(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

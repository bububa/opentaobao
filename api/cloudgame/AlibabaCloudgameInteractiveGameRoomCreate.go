package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameRoomCreate 建游戏房间
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
func AlibabaCloudgameInteractiveGameRoomCreate(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameRoomCreateAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameRoomCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

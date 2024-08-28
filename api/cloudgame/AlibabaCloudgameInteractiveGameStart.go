package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameStart 云游戏场景互动开始游戏
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
func AlibabaCloudgameInteractiveGameStart(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameStartAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameStartAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

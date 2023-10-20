package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameStart 云游戏场景互动开始游戏
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
func AlibabaCloudgameInteractiveGameStart(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameStartAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameStartAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameStartAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

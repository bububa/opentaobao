package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerStop 用户停止游戏
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
func AlibabaCloudgameInteractiveGamePlayerStop(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerKickout 踢人
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
func AlibabaCloudgameInteractiveGamePlayerKickout(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

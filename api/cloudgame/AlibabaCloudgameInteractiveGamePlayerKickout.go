package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerKickout 踢人
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
func AlibabaCloudgameInteractiveGamePlayerKickout(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

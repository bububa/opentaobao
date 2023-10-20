package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerStatusGet 获取用户状态
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
func AlibabaCloudgameInteractiveGamePlayerStatusGet(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

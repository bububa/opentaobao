package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameStatusGet 获取游戏状态
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
func AlibabaCloudgameInteractiveGameStatusGet(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameStatusGetAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameStatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

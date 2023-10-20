package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameGamepadGet 获取虚拟手柄配置
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
func AlibabaCloudgameInteractiveGameGamepadGet(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

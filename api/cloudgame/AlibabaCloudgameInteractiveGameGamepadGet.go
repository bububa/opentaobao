package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameGamepadGet 获取虚拟手柄配置
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
func AlibabaCloudgameInteractiveGameGamepadGet(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerStatusGet 获取用户状态
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
func AlibabaCloudgameInteractiveGamePlayerStatusGet(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGamePlayerStatusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

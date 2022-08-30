package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameRoomShutdown 强制关闭房间
// alibaba.cloudgame.interactive.game.room.shutdown
//
// 强制关闭房间
func AlibabaCloudgameInteractiveGameRoomShutdown(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameRoomShutdownAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

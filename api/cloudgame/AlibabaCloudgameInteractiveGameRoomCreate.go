package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameRoomCreate 建游戏房间
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
func AlibabaCloudgameInteractiveGameRoomCreate(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameRoomCreateAPIRequest, session string) (*cloudgame.AlibabaCloudgameInteractiveGameRoomCreateAPIResponse, error) {
	var resp cloudgame.AlibabaCloudgameInteractiveGameRoomCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

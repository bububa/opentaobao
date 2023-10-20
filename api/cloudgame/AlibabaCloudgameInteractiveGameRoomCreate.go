package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameroomcreate 建游戏房间
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
func Alibabacloudgameinteractivegameroomcreate(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameroomcreateAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameroomcreateAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameroomcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

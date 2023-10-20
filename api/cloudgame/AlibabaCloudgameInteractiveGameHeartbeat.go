package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameheartbeat 游戏玩家心跳
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
func Alibabacloudgameinteractivegameheartbeat(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameheartbeatAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameheartbeatAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameheartbeatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

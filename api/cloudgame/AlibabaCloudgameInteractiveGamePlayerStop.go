package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameplayerstop 用户停止游戏
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
func Alibabacloudgameinteractivegameplayerstop(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameplayerstopAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameplayerstopAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameplayerstopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

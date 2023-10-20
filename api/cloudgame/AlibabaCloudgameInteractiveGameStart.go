package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegamestart 云游戏场景互动开始游戏
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
func Alibabacloudgameinteractivegamestart(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegamestartAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegamestartAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegamestartAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

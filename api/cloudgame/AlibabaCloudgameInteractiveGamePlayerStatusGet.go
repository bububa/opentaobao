package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameplayerstatusget 获取用户状态
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
func Alibabacloudgameinteractivegameplayerstatusget(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameplayerstatusgetAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameplayerstatusgetAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameplayerstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

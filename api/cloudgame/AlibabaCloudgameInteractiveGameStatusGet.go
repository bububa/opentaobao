package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegamestatusget 获取游戏状态
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
func Alibabacloudgameinteractivegamestatusget(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegamestatusgetAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegamestatusgetAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegamestatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

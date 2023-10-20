package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameplayerkickout 踢人
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
func Alibabacloudgameinteractivegameplayerkickout(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameplayerkickoutAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameplayerkickoutAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameplayerkickoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

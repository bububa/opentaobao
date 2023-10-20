package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegameroomshutdown 强制关闭房间
// alibaba.cloudgame.interactive.game.room.shutdown
//
// 强制关闭房间
func Alibabacloudgameinteractivegameroomshutdown(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegameroomshutdownAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegameroomshutdownAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegameroomshutdownAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

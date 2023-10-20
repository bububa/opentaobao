package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacloudgameinteractivegamegamepadget 获取虚拟手柄配置
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
func Alibabacloudgameinteractivegamegamepadget(clt *core.SDKClient, req *cloudgame.AlibabacloudgameinteractivegamegamepadgetAPIRequest, session string) (*cloudgame.AlibabacloudgameinteractivegamegamepadgetAPIResponse, error) {
	var resp cloudgame.AlibabacloudgameinteractivegamegamepadgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

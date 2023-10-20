package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpprojectInitnewproject 创建新的mpproject
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
func AlibabaCgameMpMpprojectInitnewproject(clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIRequest, resp *cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

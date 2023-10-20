package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamempmpsessionsendmessagetogame 发送消息给游戏
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
func Alibabacgamempmpsessionsendmessagetogame(clt *core.SDKClient, req *cloudgame.AlibabacgamempmpsessionsendmessagetogameAPIRequest, session string) (*cloudgame.AlibabacgamempmpsessionsendmessagetogameAPIResponse, error) {
	var resp cloudgame.AlibabacgamempmpsessionsendmessagetogameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpsessionSendmessagetogame 发送消息给游戏
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
func AlibabaCgameMpMpsessionSendmessagetogame(clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIRequest, resp *cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

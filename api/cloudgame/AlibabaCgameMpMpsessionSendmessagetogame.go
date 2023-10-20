package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpsessionSendmessagetogame 发送消息给游戏
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
func AlibabaCgameMpMpsessionSendmessagetogame(clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIRequest, session string) (*cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIResponse, error) {
	var resp cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

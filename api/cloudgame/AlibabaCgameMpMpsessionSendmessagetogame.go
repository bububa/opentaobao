package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpsessionSendmessagetogame 发送消息给游戏
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
func AlibabaCgameMpMpsessionSendmessagetogame(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIRequest, resp *cloudgame.AlibabaCgameMpMpsessionSendmessagetogameAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

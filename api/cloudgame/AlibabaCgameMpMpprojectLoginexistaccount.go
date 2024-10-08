package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpprojectLoginexistaccount 登录存在账号
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
func AlibabaCgameMpMpprojectLoginexistaccount(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpprojectLoginexistaccountAPIRequest, resp *cloudgame.AlibabaCgameMpMpprojectLoginexistaccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

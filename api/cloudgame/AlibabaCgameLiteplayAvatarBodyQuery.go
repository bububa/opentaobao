package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameLiteplayAvatarBodyQuery 新氢玩Avatar脸部装扮数据查询
// alibaba.cgame.liteplay.avatar.body.query
//
// 云游戏, 新氢玩, 围观互动,自研游戏包, 需要查询Avatar虚拟形象捏脸和装扮数据, 用来初始化游戏包形象.
func AlibabaCgameLiteplayAvatarBodyQuery(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameLiteplayAvatarBodyQueryAPIRequest, resp *cloudgame.AlibabaCgameLiteplayAvatarBodyQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

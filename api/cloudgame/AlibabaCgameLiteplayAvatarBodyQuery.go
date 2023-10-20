package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameLiteplayAvatarBodyQuery 新氢玩Avatar脸部装扮数据查询
// alibaba.cgame.liteplay.avatar.body.query
//
// 云游戏, 新氢玩, 围观互动,自研游戏包, 需要查询Avatar虚拟形象捏脸和装扮数据, 用来初始化游戏包形象.
func AlibabaCgameLiteplayAvatarBodyQuery(clt *core.SDKClient, req *cloudgame.AlibabaCgameLiteplayAvatarBodyQueryAPIRequest, session string) (*cloudgame.AlibabaCgameLiteplayAvatarBodyQueryAPIResponse, error) {
	var resp cloudgame.AlibabaCgameLiteplayAvatarBodyQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

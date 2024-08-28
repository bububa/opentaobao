package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameMpMpprojectInitnewproject 创建新的mpproject
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
func AlibabaCgameMpMpprojectInitnewproject(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIRequest, resp *cloudgame.AlibabaCgameMpMpprojectInitnewprojectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

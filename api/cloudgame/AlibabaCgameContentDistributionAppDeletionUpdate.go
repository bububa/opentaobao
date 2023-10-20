package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameContentDistributionAppDeletionUpdate 游戏删除回调
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
func AlibabaCgameContentDistributionAppDeletionUpdate(clt *core.SDKClient, req *cloudgame.AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest, resp *cloudgame.AlibabaCgameContentDistributionAppDeletionUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

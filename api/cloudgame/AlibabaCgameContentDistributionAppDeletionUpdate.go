package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameContentDistributionAppDeletionUpdate 游戏删除回调
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
func AlibabaCgameContentDistributionAppDeletionUpdate(clt *core.SDKClient, req *cloudgame.AlibabaCgameContentDistributionAppDeletionUpdateAPIRequest, session string) (*cloudgame.AlibabaCgameContentDistributionAppDeletionUpdateAPIResponse, error) {
	var resp cloudgame.AlibabaCgameContentDistributionAppDeletionUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

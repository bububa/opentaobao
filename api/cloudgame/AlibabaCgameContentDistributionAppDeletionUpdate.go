package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamecontentdistributionappdeletionupdate 游戏删除回调
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
func Alibabacgamecontentdistributionappdeletionupdate(clt *core.SDKClient, req *cloudgame.AlibabacgamecontentdistributionappdeletionupdateAPIRequest, session string) (*cloudgame.AlibabacgamecontentdistributionappdeletionupdateAPIResponse, error) {
	var resp cloudgame.AlibabacgamecontentdistributionappdeletionupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

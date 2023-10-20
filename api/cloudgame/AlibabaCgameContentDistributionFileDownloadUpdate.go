package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamecontentdistributionfiledownloadupdate 文件下载回调
// alibaba.cgame.content.distribution.file.download.update
//
// 提供内容服务器更新文件下载状态的能力
func Alibabacgamecontentdistributionfiledownloadupdate(clt *core.SDKClient, req *cloudgame.AlibabacgamecontentdistributionfiledownloadupdateAPIRequest, session string) (*cloudgame.AlibabacgamecontentdistributionfiledownloadupdateAPIResponse, error) {
	var resp cloudgame.AlibabacgamecontentdistributionfiledownloadupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

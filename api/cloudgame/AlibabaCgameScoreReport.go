package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgamescorereport 云游戏战绩上传通用接口
// alibaba.cgame.score.report
//
// 阿里云游戏, CP游戏合作方通用游戏结果回传接口
func Alibabacgamescorereport(clt *core.SDKClient, req *cloudgame.AlibabacgamescorereportAPIRequest, session string) (*cloudgame.AlibabacgamescorereportAPIResponse, error) {
	var resp cloudgame.AlibabacgamescorereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

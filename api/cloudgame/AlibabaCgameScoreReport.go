package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameScoreReport 云游戏战绩上传通用接口
// alibaba.cgame.score.report
//
// 阿里云游戏, CP游戏合作方通用游戏结果回传接口
func AlibabaCgameScoreReport(clt *core.SDKClient, req *cloudgame.AlibabaCgameScoreReportAPIRequest, resp *cloudgame.AlibabaCgameScoreReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

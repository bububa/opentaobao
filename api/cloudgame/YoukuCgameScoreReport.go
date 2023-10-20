package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// YoukuCgameScoreReport 云游戏战绩上传
// youku.cgame.score.report
//
// 云游戏战绩上传API
func YoukuCgameScoreReport(clt *core.SDKClient, req *cloudgame.YoukuCgameScoreReportAPIRequest, resp *cloudgame.YoukuCgameScoreReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

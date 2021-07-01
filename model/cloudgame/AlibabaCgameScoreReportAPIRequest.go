package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCgameScoreReportAPIRequest
云游戏战绩上传通用接口 API请求
alibaba.cgame.score.report

阿里云游戏, CP游戏合作方通用游戏结果回传接口 */
type AlibabaCgameScoreReportAPIRequest struct {
	model.Params
	// 通用战绩回传数据
	_reportData *CpCallbackReportDto
}

// New

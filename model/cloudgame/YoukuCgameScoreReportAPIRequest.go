package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuCgameScoreReportAPIRequest
云游戏战绩上传 API请求
youku.cgame.score.report

云游戏战绩上传API */
type YoukuCgameScoreReportAPIRequest struct {
	model.Params
	// 战绩上传Dto
	_scoreReportDto *ScoreReportDto
}

// New

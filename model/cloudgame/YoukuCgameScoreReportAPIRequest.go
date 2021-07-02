package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuCgameScoreReportAPIRequest 云游戏战绩上传 API请求
// youku.cgame.score.report
//
// 云游戏战绩上传API
type YoukuCgameScoreReportAPIRequest struct {
	model.Params
	// 战绩上传Dto
	_scoreReportDto *ScoreReportDto
}

// NewYoukuCgameScoreReportRequest 初始化YoukuCgameScoreReportAPIRequest对象
func NewYoukuCgameScoreReportRequest() *YoukuCgameScoreReportAPIRequest {
	return &YoukuCgameScoreReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuCgameScoreReportAPIRequest) GetApiMethodName() string {
	return "youku.cgame.score.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuCgameScoreReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetScoreReportDto is ScoreReportDto Setter
// 战绩上传Dto
func (r *YoukuCgameScoreReportAPIRequest) SetScoreReportDto(_scoreReportDto *ScoreReportDto) error {
	r._scoreReportDto = _scoreReportDto
	r.Set("score_report_dto", _scoreReportDto)
	return nil
}

// GetScoreReportDto ScoreReportDto Getter
func (r YoukuCgameScoreReportAPIRequest) GetScoreReportDto() *ScoreReportDto {
	return r._scoreReportDto
}

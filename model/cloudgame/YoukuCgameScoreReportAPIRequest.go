package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukucgamescorereportAPIRequest 云游戏战绩上传 API请求
// youku.cgame.score.report
//
// 云游戏战绩上传API
type YoukucgamescorereportAPIRequest struct {
	model.Params
	// 战绩上传Dto
	_scoreReportDto *ScoreReportDto
}

// NewYoukucgamescorereportRequest 初始化YoukucgamescorereportAPIRequest对象
func NewYoukucgamescorereportRequest() *YoukucgamescorereportAPIRequest {
	return &YoukucgamescorereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukucgamescorereportAPIRequest) GetApiMethodName() string {
	return "youku.cgame.score.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukucgamescorereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukucgamescorereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScoreReportDto is ScoreReportDto Setter
// 战绩上传Dto
func (r *YoukucgamescorereportAPIRequest) SetScoreReportDto(_scoreReportDto *ScoreReportDto) error {
	r._scoreReportDto = _scoreReportDto
	r.Set("score_report_dto", _scoreReportDto)
	return nil
}

// GetScoreReportDto ScoreReportDto Getter
func (r YoukucgamescorereportAPIRequest) GetScoreReportDto() *ScoreReportDto {
	return r._scoreReportDto
}

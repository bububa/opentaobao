package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuCgameScoreReportAPIRequest) Reset() {
	r._scoreReportDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuCgameScoreReportAPIRequest) GetApiMethodName() string {
	return "youku.cgame.score.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuCgameScoreReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuCgameScoreReportAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolYoukuCgameScoreReportAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuCgameScoreReportRequest()
	},
}

// GetYoukuCgameScoreReportRequest 从 sync.Pool 获取 YoukuCgameScoreReportAPIRequest
func GetYoukuCgameScoreReportAPIRequest() *YoukuCgameScoreReportAPIRequest {
	return poolYoukuCgameScoreReportAPIRequest.Get().(*YoukuCgameScoreReportAPIRequest)
}

// ReleaseYoukuCgameScoreReportAPIRequest 将 YoukuCgameScoreReportAPIRequest 放入 sync.Pool
func ReleaseYoukuCgameScoreReportAPIRequest(v *YoukuCgameScoreReportAPIRequest) {
	v.Reset()
	poolYoukuCgameScoreReportAPIRequest.Put(v)
}

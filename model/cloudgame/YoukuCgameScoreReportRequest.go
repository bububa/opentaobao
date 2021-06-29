package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏战绩上传 API请求
youku.cgame.score.report

云游戏战绩上传API
*/
type YoukuCgameScoreReportRequest struct {
    model.Params
    // 战绩上传Dto
    _scoreReportDto   *ScoreReportDto
}

// 初始化YoukuCgameScoreReportRequest对象
func NewYoukuCgameScoreReportRequest() *YoukuCgameScoreReportRequest{
    return &YoukuCgameScoreReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuCgameScoreReportRequest) GetApiMethodName() string {
    return "youku.cgame.score.report"
}

// IRequest interface 方法, 获取API参数
func (r YoukuCgameScoreReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScoreReportDto Setter
// 战绩上传Dto
func (r *YoukuCgameScoreReportRequest) SetScoreReportDto(_scoreReportDto *ScoreReportDto) error {
    r._scoreReportDto = _scoreReportDto
    r.Set("score_report_dto", _scoreReportDto)
    return nil
}

// ScoreReportDto Getter
func (r YoukuCgameScoreReportRequest) GetScoreReportDto() *ScoreReportDto {
    return r._scoreReportDto
}

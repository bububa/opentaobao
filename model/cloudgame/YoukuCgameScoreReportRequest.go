package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏战绩上传 APIRequest
youku.cgame.score.report

云游戏战绩上传API
*/
type YoukuCgameScoreReportRequest struct {
    model.Params

    // 战绩上传Dto
    scoreReportDto   *ScoreReportDto 

}

func NewYoukuCgameScoreReportRequest() *YoukuCgameScoreReportRequest{
    return &YoukuCgameScoreReportRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuCgameScoreReportRequest) GetApiMethodName() string {
    return "youku.cgame.score.report"
}

func (r YoukuCgameScoreReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuCgameScoreReportRequest) SetScoreReportDto(scoreReportDto *ScoreReportDto) error {
    r.scoreReportDto = scoreReportDto
    r.Set("score_report_dto", scoreReportDto)
    return nil
}

func (r YoukuCgameScoreReportRequest) GetScoreReportDto() *ScoreReportDto {
    return r.scoreReportDto
}


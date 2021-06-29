package cloudgame

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏战绩上传通用接口 APIResponse
alibaba.cgame.score.report

阿里云游戏, CP游戏合作方通用游戏结果回传接口
*/
type AlibabaCgameScoreReportAPIResponse struct {
    model.CommonResponse
    AlibabaCgameScoreReportResponse
}

type AlibabaCgameScoreReportResponse struct {
    XMLName xml.Name `xml:"alibaba_cgame_score_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaCgameScoreReportResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

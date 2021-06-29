package cloudgame

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏战绩上传 APIResponse
youku.cgame.score.report

云游戏战绩上传API
*/
type YoukuCgameScoreReportAPIResponse struct {
    model.CommonResponse
    YoukuCgameScoreReportResponse
}

type YoukuCgameScoreReportResponse struct {
    XMLName xml.Name `xml:"youku_cgame_score_report_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 返回消息体
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 返回消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

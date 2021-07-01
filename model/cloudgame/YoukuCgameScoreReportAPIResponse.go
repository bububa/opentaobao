package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YoukuCgameScoreReportAPIResponse
云游戏战绩上传 API返回值
youku.cgame.score.report

云游戏战绩上传API */
type YoukuCgameScoreReportAPIResponse struct {
	model.CommonResponse
	YoukuCgameScoreReportAPIResponseModel
}

// YoukuCgameScoreReportAPIResponseModel is 云游戏战绩上传 成功返回结果
type YoukuCgameScoreReportAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_cgame_score_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回消息体
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

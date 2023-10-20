package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuCgameScoreReportAPIResponse 云游戏战绩上传 API返回值
// youku.cgame.score.report
//
// 云游戏战绩上传API
type YoukuCgameScoreReportAPIResponse struct {
	model.CommonResponse
	YoukuCgameScoreReportAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuCgameScoreReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuCgameScoreReportAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *YoukuCgameScoreReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Data = ""
	m.Message = ""
}

var poolYoukuCgameScoreReportAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuCgameScoreReportAPIResponse)
	},
}

// GetYoukuCgameScoreReportAPIResponse 从 sync.Pool 获取 YoukuCgameScoreReportAPIResponse
func GetYoukuCgameScoreReportAPIResponse() *YoukuCgameScoreReportAPIResponse {
	return poolYoukuCgameScoreReportAPIResponse.Get().(*YoukuCgameScoreReportAPIResponse)
}

// ReleaseYoukuCgameScoreReportAPIResponse 将 YoukuCgameScoreReportAPIResponse 保存到 sync.Pool
func ReleaseYoukuCgameScoreReportAPIResponse(v *YoukuCgameScoreReportAPIResponse) {
	v.Reset()
	poolYoukuCgameScoreReportAPIResponse.Put(v)
}

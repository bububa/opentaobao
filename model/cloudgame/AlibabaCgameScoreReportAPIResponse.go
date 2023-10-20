package cloudgame

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameScoreReportAPIResponse 云游戏战绩上传通用接口 API返回值
// alibaba.cgame.score.report
//
// 阿里云游戏, CP游戏合作方通用游戏结果回传接口
type AlibabaCgameScoreReportAPIResponse struct {
	model.CommonResponse
	AlibabaCgameScoreReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCgameScoreReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCgameScoreReportAPIResponseModel).Reset()
}

// AlibabaCgameScoreReportAPIResponseModel is 云游戏战绩上传通用接口 成功返回结果
type AlibabaCgameScoreReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_score_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCgameScoreReportResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCgameScoreReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCgameScoreReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCgameScoreReportAPIResponse)
	},
}

// GetAlibabaCgameScoreReportAPIResponse 从 sync.Pool 获取 AlibabaCgameScoreReportAPIResponse
func GetAlibabaCgameScoreReportAPIResponse() *AlibabaCgameScoreReportAPIResponse {
	return poolAlibabaCgameScoreReportAPIResponse.Get().(*AlibabaCgameScoreReportAPIResponse)
}

// ReleaseAlibabaCgameScoreReportAPIResponse 将 AlibabaCgameScoreReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCgameScoreReportAPIResponse(v *AlibabaCgameScoreReportAPIResponse) {
	v.Reset()
	poolAlibabaCgameScoreReportAPIResponse.Put(v)
}

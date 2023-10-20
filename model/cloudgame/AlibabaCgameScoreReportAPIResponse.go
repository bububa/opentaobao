package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamescorereportAPIResponse 云游戏战绩上传通用接口 API返回值
// alibaba.cgame.score.report
//
// 阿里云游戏, CP游戏合作方通用游戏结果回传接口
type AlibabacgamescorereportAPIResponse struct {
	model.CommonResponse
	AlibabacgamescorereportAPIResponseModel
}

// AlibabacgamescorereportAPIResponseModel is 云游戏战绩上传通用接口 成功返回结果
type AlibabacgamescorereportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_score_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabacgamescorereportResult `json:"result,omitempty" xml:"result,omitempty"`
}

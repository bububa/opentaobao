package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameLiteplayAvatarRecordReportAPIResponse Avatar形象保存地址回调 API返回值
// alibaba.cgame.liteplay.avatar.record.report
//
// 新氢玩, 围观互动Avatar捏脸, 形象地址保存回调
type AlibabaCgameLiteplayAvatarRecordReportAPIResponse struct {
	model.CommonResponse
	AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel
}

// AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel is Avatar形象保存地址回调 成功返回结果
type AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_liteplay_avatar_record_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCgameLiteplayAvatarRecordReportResult `json:"result,omitempty" xml:"result,omitempty"`
}

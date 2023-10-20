package cloudgame

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaCgameLiteplayAvatarRecordReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel).Reset()
}

// AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel is Avatar形象保存地址回调 成功返回结果
type AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_liteplay_avatar_record_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCgameLiteplayAvatarRecordReportResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCgameLiteplayAvatarRecordReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCgameLiteplayAvatarRecordReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCgameLiteplayAvatarRecordReportAPIResponse)
	},
}

// GetAlibabaCgameLiteplayAvatarRecordReportAPIResponse 从 sync.Pool 获取 AlibabaCgameLiteplayAvatarRecordReportAPIResponse
func GetAlibabaCgameLiteplayAvatarRecordReportAPIResponse() *AlibabaCgameLiteplayAvatarRecordReportAPIResponse {
	return poolAlibabaCgameLiteplayAvatarRecordReportAPIResponse.Get().(*AlibabaCgameLiteplayAvatarRecordReportAPIResponse)
}

// ReleaseAlibabaCgameLiteplayAvatarRecordReportAPIResponse 将 AlibabaCgameLiteplayAvatarRecordReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCgameLiteplayAvatarRecordReportAPIResponse(v *AlibabaCgameLiteplayAvatarRecordReportAPIResponse) {
	v.Reset()
	poolAlibabaCgameLiteplayAvatarRecordReportAPIResponse.Put(v)
}

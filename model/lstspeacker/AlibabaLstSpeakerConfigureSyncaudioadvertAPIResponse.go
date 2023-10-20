package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse 同步广告 API返回值
// alibaba.lst.speaker.configure.syncaudioadvert
//
// 如意音箱广告同步
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel).Reset()
}

// AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel is 同步广告 成功返回结果
type AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_syncaudioadvert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErroMessage string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
	// 错误码
	ErroCode string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 执行结果标识
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErroMessage = ""
	m.ErroCode = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse)
	},
}

// GetAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse
func GetAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse() *AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse {
	return poolAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse.Get().(*AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse)
}

// ReleaseAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse 将 AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse(v *AlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureSyncaudioadvertAPIResponse.Put(v)
}

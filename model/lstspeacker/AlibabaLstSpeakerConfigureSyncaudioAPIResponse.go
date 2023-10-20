package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureSyncaudioAPIResponse 音频同步 API返回值
// alibaba.lst.speaker.configure.syncaudio
//
// 音频同步
type AlibabaLstSpeakerConfigureSyncaudioAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerConfigureSyncaudioAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureSyncaudioAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerConfigureSyncaudioAPIResponseModel).Reset()
}

// AlibabaLstSpeakerConfigureSyncaudioAPIResponseModel is 音频同步 成功返回结果
type AlibabaLstSpeakerConfigureSyncaudioAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_syncaudio_response"`
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
func (m *AlibabaLstSpeakerConfigureSyncaudioAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErroMessage = ""
	m.ErroCode = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstSpeakerConfigureSyncaudioAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerConfigureSyncaudioAPIResponse)
	},
}

// GetAlibabaLstSpeakerConfigureSyncaudioAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerConfigureSyncaudioAPIResponse
func GetAlibabaLstSpeakerConfigureSyncaudioAPIResponse() *AlibabaLstSpeakerConfigureSyncaudioAPIResponse {
	return poolAlibabaLstSpeakerConfigureSyncaudioAPIResponse.Get().(*AlibabaLstSpeakerConfigureSyncaudioAPIResponse)
}

// ReleaseAlibabaLstSpeakerConfigureSyncaudioAPIResponse 将 AlibabaLstSpeakerConfigureSyncaudioAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureSyncaudioAPIResponse(v *AlibabaLstSpeakerConfigureSyncaudioAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureSyncaudioAPIResponse.Put(v)
}

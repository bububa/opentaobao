package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse 音箱音量调节 API返回值
// alibaba.lst.speaker.configure.adjustvolume
//
// 音箱音量调节
type AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerConfigureAdjustvolumeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerConfigureAdjustvolumeAPIResponseModel).Reset()
}

// AlibabaLstSpeakerConfigureAdjustvolumeAPIResponseModel is 音箱音量调节 成功返回结果
type AlibabaLstSpeakerConfigureAdjustvolumeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_adjustvolume_response"`
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
func (m *AlibabaLstSpeakerConfigureAdjustvolumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErroMessage = ""
	m.ErroCode = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse)
	},
}

// GetAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse
func GetAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse() *AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse {
	return poolAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse.Get().(*AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse)
}

// ReleaseAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse 将 AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse(v *AlibabaLstSpeakerConfigureAdjustvolumeAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureAdjustvolumeAPIResponse.Put(v)
}

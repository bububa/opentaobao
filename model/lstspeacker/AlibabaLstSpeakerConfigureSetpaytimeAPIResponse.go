package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureSetpaytimeAPIResponse 音箱播放配置 API返回值
// alibaba.lst.speaker.configure.setpaytime
//
// 音箱播放配置
type AlibabaLstSpeakerConfigureSetpaytimeAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerConfigureSetpaytimeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureSetpaytimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerConfigureSetpaytimeAPIResponseModel).Reset()
}

// AlibabaLstSpeakerConfigureSetpaytimeAPIResponseModel is 音箱播放配置 成功返回结果
type AlibabaLstSpeakerConfigureSetpaytimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_setpaytime_response"`
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
func (m *AlibabaLstSpeakerConfigureSetpaytimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErroMessage = ""
	m.ErroCode = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstSpeakerConfigureSetpaytimeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerConfigureSetpaytimeAPIResponse)
	},
}

// GetAlibabaLstSpeakerConfigureSetpaytimeAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerConfigureSetpaytimeAPIResponse
func GetAlibabaLstSpeakerConfigureSetpaytimeAPIResponse() *AlibabaLstSpeakerConfigureSetpaytimeAPIResponse {
	return poolAlibabaLstSpeakerConfigureSetpaytimeAPIResponse.Get().(*AlibabaLstSpeakerConfigureSetpaytimeAPIResponse)
}

// ReleaseAlibabaLstSpeakerConfigureSetpaytimeAPIResponse 将 AlibabaLstSpeakerConfigureSetpaytimeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureSetpaytimeAPIResponse(v *AlibabaLstSpeakerConfigureSetpaytimeAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureSetpaytimeAPIResponse.Put(v)
}

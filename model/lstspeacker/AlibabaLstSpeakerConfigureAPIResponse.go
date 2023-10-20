package lstspeacker

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstSpeakerConfigureAPIResponse 零售通音箱配置通用泛化调用接口 API返回值
// alibaba.lst.speaker.configure
//
// 零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
type AlibabaLstSpeakerConfigureAPIResponse struct {
	model.CommonResponse
	AlibabaLstSpeakerConfigureAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstSpeakerConfigureAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstSpeakerConfigureAPIResponseModel).Reset()
}

// AlibabaLstSpeakerConfigureAPIResponseModel is 零售通音箱配置通用泛化调用接口 成功返回结果
type AlibabaLstSpeakerConfigureAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_speaker_configure_response"`
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
func (m *AlibabaLstSpeakerConfigureAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErroMessage = ""
	m.ErroCode = ""
	m.Succ = false
	m.Module = false
}

var poolAlibabaLstSpeakerConfigureAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerConfigureAPIResponse)
	},
}

// GetAlibabaLstSpeakerConfigureAPIResponse 从 sync.Pool 获取 AlibabaLstSpeakerConfigureAPIResponse
func GetAlibabaLstSpeakerConfigureAPIResponse() *AlibabaLstSpeakerConfigureAPIResponse {
	return poolAlibabaLstSpeakerConfigureAPIResponse.Get().(*AlibabaLstSpeakerConfigureAPIResponse)
}

// ReleaseAlibabaLstSpeakerConfigureAPIResponse 将 AlibabaLstSpeakerConfigureAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstSpeakerConfigureAPIResponse(v *AlibabaLstSpeakerConfigureAPIResponse) {
	v.Reset()
	poolAlibabaLstSpeakerConfigureAPIResponse.Put(v)
}

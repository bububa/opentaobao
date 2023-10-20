package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsUserSpeechGuideAPIResponse 引导语推荐接口 API返回值
// alibaba.ailabs.user.speech.guide
//
// 根据用户的语音query，返回相应的引导语推荐
type AlibabaAilabsUserSpeechGuideAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsUserSpeechGuideAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsUserSpeechGuideAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsUserSpeechGuideAPIResponseModel).Reset()
}

// AlibabaAilabsUserSpeechGuideAPIResponseModel is 引导语推荐接口 成功返回结果
type AlibabaAilabsUserSpeechGuideAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_user_speech_guide_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsUserSpeechGuideResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsUserSpeechGuideAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsUserSpeechGuideAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsUserSpeechGuideAPIResponse)
	},
}

// GetAlibabaAilabsUserSpeechGuideAPIResponse 从 sync.Pool 获取 AlibabaAilabsUserSpeechGuideAPIResponse
func GetAlibabaAilabsUserSpeechGuideAPIResponse() *AlibabaAilabsUserSpeechGuideAPIResponse {
	return poolAlibabaAilabsUserSpeechGuideAPIResponse.Get().(*AlibabaAilabsUserSpeechGuideAPIResponse)
}

// ReleaseAlibabaAilabsUserSpeechGuideAPIResponse 将 AlibabaAilabsUserSpeechGuideAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsUserSpeechGuideAPIResponse(v *AlibabaAilabsUserSpeechGuideAPIResponse) {
	v.Reset()
	poolAlibabaAilabsUserSpeechGuideAPIResponse.Put(v)
}

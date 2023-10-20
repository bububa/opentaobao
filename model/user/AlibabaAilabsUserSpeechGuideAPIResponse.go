package user

import (
	"encoding/xml"

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

// AlibabaAilabsUserSpeechGuideAPIResponseModel is 引导语推荐接口 成功返回结果
type AlibabaAilabsUserSpeechGuideAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_user_speech_guide_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsUserSpeechGuideResult `json:"result,omitempty" xml:"result,omitempty"`
}

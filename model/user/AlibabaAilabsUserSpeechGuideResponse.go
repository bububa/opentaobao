package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
引导语推荐接口 APIResponse
alibaba.ailabs.user.speech.guide

根据用户的语音query，返回相应的引导语推荐
*/
type AlibabaAilabsUserSpeechGuideAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAilabsUserSpeechGuideResponse `json:"alibaba_ailabs_user_speech_guide_response,omitempty"`
}

type AlibabaAilabsUserSpeechGuideResponse struct {

    // 接口返回model
    Result   *AlibabaAilabsUserSpeechGuideResult `json:"result,omitempty"`

}

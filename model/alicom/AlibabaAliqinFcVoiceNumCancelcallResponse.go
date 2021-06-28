package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消呼叫 APIResponse
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
*/
type AlibabaAliqinFcVoiceNumCancelcallAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcVoiceNumCancelcallResponse `json:"alibaba_aliqin_fc_voice_num_cancelcall_response,omitempty"` 
    AlibabaAliqinFcVoiceNumCancelcallResponse
}

/* model for simplify = false
type AlibabaAliqinFcVoiceNumCancelcallResponse struct {

    // result
    
    Result  *struct {
        AlibabaAliqinFcVoiceNumCancelcallBizResult  *AlibabaAliqinFcVoiceNumCancelcallBizResult `json:"alibaba_aliqin_fc_voice_num_cancelcall_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcVoiceNumCancelcallResponse struct {

    // result
    Result   *AlibabaAliqinFcVoiceNumCancelcallBizResult `json:"result,omitempty"`

}

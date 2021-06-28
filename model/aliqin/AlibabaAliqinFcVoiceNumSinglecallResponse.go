package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
语音通知 APIResponse
alibaba.aliqin.fc.voice.num.singlecall

向指定手机号码发起单向呼叫，播放指定的语音文件内容。使用前需要在阿里大于管理中心添加去电显示号码与语音文件。
*/
type AlibabaAliqinFcVoiceNumSinglecallAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcVoiceNumSinglecallResponse `json:"alibaba_aliqin_fc_voice_num_singlecall_response,omitempty"` 
    AlibabaAliqinFcVoiceNumSinglecallResponse
}

/* model for simplify = false
type AlibabaAliqinFcVoiceNumSinglecallResponse struct {

    // 返回值
    
    Result  *struct {
        AlibabaAliqinFcVoiceNumSinglecallBizResult  *AlibabaAliqinFcVoiceNumSinglecallBizResult `json:"alibaba_aliqin_fc_voice_num_singlecall_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcVoiceNumSinglecallResponse struct {

    // 返回值
    Result   *AlibabaAliqinFcVoiceNumSinglecallBizResult `json:"result,omitempty"`

}

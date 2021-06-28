package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
多方通话 APIResponse
alibaba.aliqin.fc.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
type AlibabaAliqinFcVoiceNumDoublecallAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcVoiceNumDoublecallResponse `json:"alibaba_aliqin_fc_voice_num_doublecall_response,omitempty"` 
    AlibabaAliqinFcVoiceNumDoublecallResponse
}

/* model for simplify = false
type AlibabaAliqinFcVoiceNumDoublecallResponse struct {

    // 接口返回参数
    
    Result  *struct {
        AlibabaAliqinFcVoiceNumDoublecallBizResult  *AlibabaAliqinFcVoiceNumDoublecallBizResult `json:"alibaba_aliqin_fc_voice_num_doublecall_biz_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinFcVoiceNumDoublecallResponse struct {

    // 接口返回参数
    Result   *AlibabaAliqinFcVoiceNumDoublecallBizResult `json:"result,omitempty"`

}

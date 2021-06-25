package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔语音双呼接口 APIResponse
alibaba.aliqin.ta.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
type AlibabaAliqinTaVoiceNumDoublecallAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinTaVoiceNumDoublecallResponse `json:"alibaba_aliqin_ta_voice_num_doublecall_response,omitempty"`
}

type AlibabaAliqinTaVoiceNumDoublecallResponse struct {

    // 接口返回参数
    Result   *AlibabaAliqinTaVoiceNumDoublecallBizResult `json:"result,omitempty"`

}

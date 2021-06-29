package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
多方通话 APIResponse
alibaba.aliqin.fc.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
type AlibabaAliqinFcVoiceNumDoublecallAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcVoiceNumDoublecallResponse
}

type AlibabaAliqinFcVoiceNumDoublecallResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_doublecall_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回参数
    
    Result   *AlibabaAliqinFcVoiceNumDoublecallBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

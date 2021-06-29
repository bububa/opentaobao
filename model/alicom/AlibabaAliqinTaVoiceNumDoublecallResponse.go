package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔语音双呼接口 API返回值 
alibaba.aliqin.ta.voice.num.doublecall

根据传入的主叫号码与被叫号码，由系统依次向主叫号码与被叫号码发起呼叫，双方都应答后，开始一对一通话并开始计费。使用前需要在阿里大于管理中心添加呼叫双方的显示号码。
*/
type AlibabaAliqinTaVoiceNumDoublecallAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinTaVoiceNumDoublecallResponse
}

// 聚石塔语音双呼接口 成功返回结果
type AlibabaAliqinTaVoiceNumDoublecallResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_ta_voice_num_doublecall_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回参数
    Result   *AlibabaAliqinTaVoiceNumDoublecallBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

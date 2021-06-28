package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消呼叫 APIResponse
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
*/
type AlibabaAliqinFcVoiceNumCancelcallAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcVoiceNumCancelcallResponse
}

type AlibabaAliqinFcVoiceNumCancelcallResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_num_cancelcall_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcVoiceNumCancelcallBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

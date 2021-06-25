package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消呼叫 APIRequest
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
*/
type AlibabaAliqinFcVoiceNumCancelcallRequest struct {
    model.Params

    // 呼叫唯一id
    callId   string 

}

func NewAlibabaAliqinFcVoiceNumCancelcallRequest() *AlibabaAliqinFcVoiceNumCancelcallRequest{
    return &AlibabaAliqinFcVoiceNumCancelcallRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcVoiceNumCancelcallRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.num.cancelcall"
}

func (r AlibabaAliqinFcVoiceNumCancelcallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcVoiceNumCancelcallRequest) SetCallId(callId string) error {
    r.callId = callId
    r.Set("call_id", callId)
    return nil
}

func (r AlibabaAliqinFcVoiceNumCancelcallRequest) GetCallId() string {
    return r.callId
}


package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
录音文件下载 APIRequest
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
type AlibabaAliqinFcVoiceRecordGeturlRequest struct {
    model.Params

    // 一次通话的唯一标识id
    callId   string 

}

func NewAlibabaAliqinFcVoiceRecordGeturlRequest() *AlibabaAliqinFcVoiceRecordGeturlRequest{
    return &AlibabaAliqinFcVoiceRecordGeturlRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.record.geturl"
}

func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcVoiceRecordGeturlRequest) SetCallId(callId string) error {
    r.callId = callId
    r.Set("call_id", callId)
    return nil
}

func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetCallId() string {
    return r.callId
}


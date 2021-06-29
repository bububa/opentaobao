package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
录音文件下载 API请求
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
type AlibabaAliqinFcVoiceRecordGeturlRequest struct {
    model.Params
    // 一次通话的唯一标识id
    _callId   string
}

// 初始化AlibabaAliqinFcVoiceRecordGeturlRequest对象
func NewAlibabaAliqinFcVoiceRecordGeturlRequest() *AlibabaAliqinFcVoiceRecordGeturlRequest{
    return &AlibabaAliqinFcVoiceRecordGeturlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.record.geturl"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallId Setter
// 一次通话的唯一标识id
func (r *AlibabaAliqinFcVoiceRecordGeturlRequest) SetCallId(_callId string) error {
    r._callId = _callId
    r.Set("call_id", _callId)
    return nil
}

// CallId Getter
func (r AlibabaAliqinFcVoiceRecordGeturlRequest) GetCallId() string {
    return r._callId
}

package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceRecordGeturlAPIRequest 录音文件下载 API请求
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
type AlibabaAliqinFcVoiceRecordGeturlAPIRequest struct {
	model.Params
	// 一次通话的唯一标识id
	_callId string
}

// NewAlibabaAliqinFcVoiceRecordGeturlRequest 初始化AlibabaAliqinFcVoiceRecordGeturlAPIRequest对象
func NewAlibabaAliqinFcVoiceRecordGeturlRequest() *AlibabaAliqinFcVoiceRecordGeturlAPIRequest {
	return &AlibabaAliqinFcVoiceRecordGeturlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.record.geturl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallId Setter
// 一次通话的唯一标识id
func (r *AlibabaAliqinFcVoiceRecordGeturlAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// Get CallId Getter
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetCallId() string {
	return r._callId
}

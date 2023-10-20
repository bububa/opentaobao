package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcVoiceRecordGeturlAPIRequest) Reset() {
	r._callId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.record.geturl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallId is CallId Setter
// 一次通话的唯一标识id
func (r *AlibabaAliqinFcVoiceRecordGeturlAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaAliqinFcVoiceRecordGeturlAPIRequest) GetCallId() string {
	return r._callId
}

var poolAlibabaAliqinFcVoiceRecordGeturlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcVoiceRecordGeturlRequest()
	},
}

// GetAlibabaAliqinFcVoiceRecordGeturlRequest 从 sync.Pool 获取 AlibabaAliqinFcVoiceRecordGeturlAPIRequest
func GetAlibabaAliqinFcVoiceRecordGeturlAPIRequest() *AlibabaAliqinFcVoiceRecordGeturlAPIRequest {
	return poolAlibabaAliqinFcVoiceRecordGeturlAPIRequest.Get().(*AlibabaAliqinFcVoiceRecordGeturlAPIRequest)
}

// ReleaseAlibabaAliqinFcVoiceRecordGeturlAPIRequest 将 AlibabaAliqinFcVoiceRecordGeturlAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcVoiceRecordGeturlAPIRequest(v *AlibabaAliqinFcVoiceRecordGeturlAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcVoiceRecordGeturlAPIRequest.Put(v)
}

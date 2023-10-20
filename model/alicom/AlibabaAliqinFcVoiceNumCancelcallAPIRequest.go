package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceNumCancelcallAPIRequest 取消呼叫 API请求
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
type AlibabaAliqinFcVoiceNumCancelcallAPIRequest struct {
	model.Params
	// 呼叫唯一id
	_callId string
}

// NewAlibabaAliqinFcVoiceNumCancelcallRequest 初始化AlibabaAliqinFcVoiceNumCancelcallAPIRequest对象
func NewAlibabaAliqinFcVoiceNumCancelcallRequest() *AlibabaAliqinFcVoiceNumCancelcallAPIRequest {
	return &AlibabaAliqinFcVoiceNumCancelcallAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcVoiceNumCancelcallAPIRequest) Reset() {
	r._callId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.num.cancelcall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallId is CallId Setter
// 呼叫唯一id
func (r *AlibabaAliqinFcVoiceNumCancelcallAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetCallId() string {
	return r._callId
}

var poolAlibabaAliqinFcVoiceNumCancelcallAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcVoiceNumCancelcallRequest()
	},
}

// GetAlibabaAliqinFcVoiceNumCancelcallRequest 从 sync.Pool 获取 AlibabaAliqinFcVoiceNumCancelcallAPIRequest
func GetAlibabaAliqinFcVoiceNumCancelcallAPIRequest() *AlibabaAliqinFcVoiceNumCancelcallAPIRequest {
	return poolAlibabaAliqinFcVoiceNumCancelcallAPIRequest.Get().(*AlibabaAliqinFcVoiceNumCancelcallAPIRequest)
}

// ReleaseAlibabaAliqinFcVoiceNumCancelcallAPIRequest 将 AlibabaAliqinFcVoiceNumCancelcallAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcVoiceNumCancelcallAPIRequest(v *AlibabaAliqinFcVoiceNumCancelcallAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcVoiceNumCancelcallAPIRequest.Put(v)
}

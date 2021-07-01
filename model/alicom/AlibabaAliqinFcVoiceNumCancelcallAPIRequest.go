package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcVoiceNumCancelcallAPIRequest
取消呼叫 API请求
alibaba.aliqin.fc.voice.num.cancelcall

当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话 */
type AlibabaAliqinFcVoiceNumCancelcallAPIRequest struct {
	model.Params
	// 呼叫唯一id
	_callId string
}

// NewAlibabaAliqinFcVoiceNumCancelcallRequest 初始化AlibabaAliqinFcVoiceNumCancelcallAPIRequest对象
func NewAlibabaAliqinFcVoiceNumCancelcallRequest() *AlibabaAliqinFcVoiceNumCancelcallAPIRequest {
	return &AlibabaAliqinFcVoiceNumCancelcallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.num.cancelcall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallId Setter
// 呼叫唯一id
func (r *AlibabaAliqinFcVoiceNumCancelcallAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// Get CallId Getter
func (r AlibabaAliqinFcVoiceNumCancelcallAPIRequest) GetCallId() string {
	return r._callId
}

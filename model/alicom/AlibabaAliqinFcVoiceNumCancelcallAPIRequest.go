package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcvoicenumcancelcallAPIRequest 取消呼叫 API请求
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
type AlibabaaliqinfcvoicenumcancelcallAPIRequest struct {
	model.Params
	// 呼叫唯一id
	_callId string
}

// NewAlibabaaliqinfcvoicenumcancelcallRequest 初始化AlibabaaliqinfcvoicenumcancelcallAPIRequest对象
func NewAlibabaaliqinfcvoicenumcancelcallRequest() *AlibabaaliqinfcvoicenumcancelcallAPIRequest {
	return &AlibabaaliqinfcvoicenumcancelcallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfcvoicenumcancelcallAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.num.cancelcall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfcvoicenumcancelcallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfcvoicenumcancelcallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallId is CallId Setter
// 呼叫唯一id
func (r *AlibabaaliqinfcvoicenumcancelcallAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaaliqinfcvoicenumcancelcallAPIRequest) GetCallId() string {
	return r._callId
}

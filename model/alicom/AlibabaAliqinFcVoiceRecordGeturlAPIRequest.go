package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcvoicerecordgeturlAPIRequest 录音文件下载 API请求
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
type AlibabaaliqinfcvoicerecordgeturlAPIRequest struct {
	model.Params
	// 一次通话的唯一标识id
	_callId string
}

// NewAlibabaaliqinfcvoicerecordgeturlRequest 初始化AlibabaaliqinfcvoicerecordgeturlAPIRequest对象
func NewAlibabaaliqinfcvoicerecordgeturlRequest() *AlibabaaliqinfcvoicerecordgeturlAPIRequest {
	return &AlibabaaliqinfcvoicerecordgeturlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfcvoicerecordgeturlAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.record.geturl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfcvoicerecordgeturlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfcvoicerecordgeturlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallId is CallId Setter
// 一次通话的唯一标识id
func (r *AlibabaaliqinfcvoicerecordgeturlAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaaliqinfcvoicerecordgeturlAPIRequest) GetCallId() string {
	return r._callId
}

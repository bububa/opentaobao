package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinaxbvendorcallcontrolAPIRequest 转呼控制接口 API请求
// alibaba.aliqin.axb.vendor.call.control
//
// 转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
type AlibabaaliqinaxbvendorcallcontrolAPIRequest struct {
	model.Params
	// 转接控制接口request对象
	_startCallRequest *StartCallRequest
}

// NewAlibabaaliqinaxbvendorcallcontrolRequest 初始化AlibabaaliqinaxbvendorcallcontrolAPIRequest对象
func NewAlibabaaliqinaxbvendorcallcontrolRequest() *AlibabaaliqinaxbvendorcallcontrolAPIRequest {
	return &AlibabaaliqinaxbvendorcallcontrolAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinaxbvendorcallcontrolAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.call.control"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinaxbvendorcallcontrolAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinaxbvendorcallcontrolAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartCallRequest is StartCallRequest Setter
// 转接控制接口request对象
func (r *AlibabaaliqinaxbvendorcallcontrolAPIRequest) SetStartCallRequest(_startCallRequest *StartCallRequest) error {
	r._startCallRequest = _startCallRequest
	r.Set("start_call_request", _startCallRequest)
	return nil
}

// GetStartCallRequest StartCallRequest Getter
func (r AlibabaaliqinaxbvendorcallcontrolAPIRequest) GetStartCallRequest() *StartCallRequest {
	return r._startCallRequest
}

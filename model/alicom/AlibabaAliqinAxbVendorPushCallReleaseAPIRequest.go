package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinaxbvendorpushcallreleaseAPIRequest 供应商推送通话结束事件 API请求
// alibaba.aliqin.axb.vendor.push.call.release
//
// 通话结束挂断的时候，供应商推送通话结束事件给阿里侧
type AlibabaaliqinaxbvendorpushcallreleaseAPIRequest struct {
	model.Params
	// end_call_request
	_endCallRequest *EndCallRequest
}

// NewAlibabaaliqinaxbvendorpushcallreleaseRequest 初始化AlibabaaliqinaxbvendorpushcallreleaseAPIRequest对象
func NewAlibabaaliqinaxbvendorpushcallreleaseRequest() *AlibabaaliqinaxbvendorpushcallreleaseAPIRequest {
	return &AlibabaaliqinaxbvendorpushcallreleaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinaxbvendorpushcallreleaseAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.push.call.release"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinaxbvendorpushcallreleaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinaxbvendorpushcallreleaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndCallRequest is EndCallRequest Setter
// end_call_request
func (r *AlibabaaliqinaxbvendorpushcallreleaseAPIRequest) SetEndCallRequest(_endCallRequest *EndCallRequest) error {
	r._endCallRequest = _endCallRequest
	r.Set("end_call_request", _endCallRequest)
	return nil
}

// GetEndCallRequest EndCallRequest Getter
func (r AlibabaaliqinaxbvendorpushcallreleaseAPIRequest) GetEndCallRequest() *EndCallRequest {
	return r._endCallRequest
}

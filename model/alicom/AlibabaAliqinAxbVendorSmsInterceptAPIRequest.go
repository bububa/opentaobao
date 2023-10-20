package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinaxbvendorsmsinterceptAPIRequest AXB短信托收推送接口 API请求
// alibaba.aliqin.axb.vendor.sms.intercept
//
// 用于给供应商推送需要托收的短信
type AlibabaaliqinaxbvendorsmsinterceptAPIRequest struct {
	model.Params
	// 短信托收结构体
	_smsInterceptRequest *SmsInterceptRequest
}

// NewAlibabaaliqinaxbvendorsmsinterceptRequest 初始化AlibabaaliqinaxbvendorsmsinterceptAPIRequest对象
func NewAlibabaaliqinaxbvendorsmsinterceptRequest() *AlibabaaliqinaxbvendorsmsinterceptAPIRequest {
	return &AlibabaaliqinaxbvendorsmsinterceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinaxbvendorsmsinterceptAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.sms.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinaxbvendorsmsinterceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinaxbvendorsmsinterceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsInterceptRequest is SmsInterceptRequest Setter
// 短信托收结构体
func (r *AlibabaaliqinaxbvendorsmsinterceptAPIRequest) SetSmsInterceptRequest(_smsInterceptRequest *SmsInterceptRequest) error {
	r._smsInterceptRequest = _smsInterceptRequest
	r.Set("sms_intercept_request", _smsInterceptRequest)
	return nil
}

// GetSmsInterceptRequest SmsInterceptRequest Getter
func (r AlibabaaliqinaxbvendorsmsinterceptAPIRequest) GetSmsInterceptRequest() *SmsInterceptRequest {
	return r._smsInterceptRequest
}

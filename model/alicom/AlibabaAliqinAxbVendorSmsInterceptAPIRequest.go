package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorSmsInterceptAPIRequest AXB短信托收推送接口 API请求
// alibaba.aliqin.axb.vendor.sms.intercept
//
// 用于给供应商推送需要托收的短信
type AlibabaAliqinAxbVendorSmsInterceptAPIRequest struct {
	model.Params
	// 短信托收结构体
	_smsInterceptRequest *SmsInterceptRequest
}

// NewAlibabaAliqinAxbVendorSmsInterceptRequest 初始化AlibabaAliqinAxbVendorSmsInterceptAPIRequest对象
func NewAlibabaAliqinAxbVendorSmsInterceptRequest() *AlibabaAliqinAxbVendorSmsInterceptAPIRequest {
	return &AlibabaAliqinAxbVendorSmsInterceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.sms.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSmsInterceptRequest is SmsInterceptRequest Setter
// 短信托收结构体
func (r *AlibabaAliqinAxbVendorSmsInterceptAPIRequest) SetSmsInterceptRequest(_smsInterceptRequest *SmsInterceptRequest) error {
	r._smsInterceptRequest = _smsInterceptRequest
	r.Set("sms_intercept_request", _smsInterceptRequest)
	return nil
}

// GetSmsInterceptRequest SmsInterceptRequest Getter
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetSmsInterceptRequest() *SmsInterceptRequest {
	return r._smsInterceptRequest
}

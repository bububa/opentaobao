package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinAxbVendorSmsInterceptAPIRequest) Reset() {
	r._smsInterceptRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.sms.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinAxbVendorSmsInterceptAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAliqinAxbVendorSmsInterceptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinAxbVendorSmsInterceptRequest()
	},
}

// GetAlibabaAliqinAxbVendorSmsInterceptRequest 从 sync.Pool 获取 AlibabaAliqinAxbVendorSmsInterceptAPIRequest
func GetAlibabaAliqinAxbVendorSmsInterceptAPIRequest() *AlibabaAliqinAxbVendorSmsInterceptAPIRequest {
	return poolAlibabaAliqinAxbVendorSmsInterceptAPIRequest.Get().(*AlibabaAliqinAxbVendorSmsInterceptAPIRequest)
}

// ReleaseAlibabaAliqinAxbVendorSmsInterceptAPIRequest 将 AlibabaAliqinAxbVendorSmsInterceptAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinAxbVendorSmsInterceptAPIRequest(v *AlibabaAliqinAxbVendorSmsInterceptAPIRequest) {
	v.Reset()
	poolAlibabaAliqinAxbVendorSmsInterceptAPIRequest.Put(v)
}

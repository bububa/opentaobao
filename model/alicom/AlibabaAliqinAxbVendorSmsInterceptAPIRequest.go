package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorSmsInterceptAPIRequest
AXB短信托收推送接口 API请求
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信 */
type AlibabaAliqinAxbVendorSmsInterceptAPIRequest struct {
	model.Params
	// 短信托收结构体
	_smsInterceptRequest *SmsInterceptRequest
}

// New

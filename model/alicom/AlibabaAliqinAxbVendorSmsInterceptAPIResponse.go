package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorSmsInterceptAPIResponse
AXB短信托收推送接口 API返回值
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信 */
type AlibabaAliqinAxbVendorSmsInterceptAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorSmsInterceptAPIResponseModel
}

// AlibabaAliqinAxbVendorSmsInterceptAPIResponseModel is AXB短信托收推送接口 成功返回结果
type AlibabaAliqinAxbVendorSmsInterceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_sms_intercept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结构体
	Result *AlibabaAliqinAxbVendorSmsInterceptResponse `json:"result,omitempty" xml:"result,omitempty"`
}

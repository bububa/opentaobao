package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AXB短信托收推送接口 APIResponse
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信
*/
type AlibabaAliqinAxbVendorSmsInterceptAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_axb_vendor_sms_intercept_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应结构体
    
    Result   *Response `json:"result,omitempty" xml:"
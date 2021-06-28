package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AXB短信托收推送接口 APIResponse
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信
*/
type AlibabaAliqinAxbVendorSmsInterceptAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinAxbVendorSmsInterceptResponse `json:"alibaba_aliqin_axb_vendor_sms_intercept_response,omitempty"` 
    AlibabaAliqinAxbVendorSmsInterceptResponse
}

/* model for simplify = false
type AlibabaAliqinAxbVendorSmsInterceptResponse struct {

    // 响应结构体
    
    Result  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinAxbVendorSmsInterceptResponse struct {

    // 响应结构体
    Result   *Response `json:"result,omitempty"`

}

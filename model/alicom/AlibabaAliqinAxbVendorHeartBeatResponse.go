package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商心跳上报接口 APIResponse
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息
*/
type AlibabaAliqinAxbVendorHeartBeatAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinAxbVendorHeartBeatResponse `json:"alibaba_aliqin_axb_vendor_heart_beat_response,omitempty"` 
    AlibabaAliqinAxbVendorHeartBeatResponse
}

/* model for simplify = false
type AlibabaAliqinAxbVendorHeartBeatResponse struct {

    // result
    
    Result  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAliqinAxbVendorHeartBeatResponse struct {

    // result
    Result   *Response `json:"result,omitempty"`

}

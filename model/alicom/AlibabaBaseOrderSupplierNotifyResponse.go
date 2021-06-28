package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信运营商信息回传 APIResponse
alibaba.base.order.supplier.notify

接收阿里通信流量运营商信息回传
*/
type AlibabaBaseOrderSupplierNotifyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBaseOrderSupplierNotifyResponse `json:"alibaba_base_order_supplier_notify_response,omitempty"` 
    AlibabaBaseOrderSupplierNotifyResponse
}

/* model for simplify = false
type AlibabaBaseOrderSupplierNotifyResponse struct {

    // result
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaBaseOrderSupplierNotifyResponse struct {

    // result
    Result   *CommonResult `json:"result,omitempty"`

}

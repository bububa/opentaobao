package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信运营商信息回传 APIResponse
alibaba.base.order.supplier.notify

接收阿里通信流量运营商信息回传
*/
type AlibabaBaseOrderSupplierNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaBaseOrderSupplierNotifyResponse
}

type AlibabaBaseOrderSupplierNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_base_order_supplier_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

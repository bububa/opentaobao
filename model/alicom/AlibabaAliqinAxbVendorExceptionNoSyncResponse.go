package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中心化供应商异常号码状态同步接口 APIResponse
alibaba.aliqin.axb.vendor.exception.no.sync

用于中心化供应商同步异常号码
*/
type AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinAxbVendorExceptionNoSyncResponse
}

type AlibabaAliqinAxbVendorExceptionNoSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_exception_no_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}

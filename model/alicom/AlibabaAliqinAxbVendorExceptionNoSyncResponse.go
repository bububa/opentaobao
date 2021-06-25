package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
中心化供应商异常号码状态同步接口 APIResponse
alibaba.aliqin.axb.vendor.exception.no.sync

用于中心化供应商同步异常号码
*/
type AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinAxbVendorExceptionNoSyncResponse `json:"alibaba_aliqin_axb_vendor_exception_no_sync_response,omitempty"`
}

type AlibabaAliqinAxbVendorExceptionNoSyncResponse struct {

    // result
    Result   *Response `json:"result,omitempty"`

}

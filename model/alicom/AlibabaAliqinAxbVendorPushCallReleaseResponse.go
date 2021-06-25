package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商推送通话结束事件 APIResponse
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinAxbVendorPushCallReleaseResponse `json:"alibaba_aliqin_axb_vendor_push_call_release_response,omitempty"`
}

type AlibabaAliqinAxbVendorPushCallReleaseResponse struct {

    // result
    Result   *Response `json:"result,omitempty"`

}

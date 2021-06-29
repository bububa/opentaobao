package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商推送通话结束事件 APIResponse
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinAxbVendorPushCallReleaseResponse
}

type AlibabaAliqinAxbVendorPushCallReleaseResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_push_call_release_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}

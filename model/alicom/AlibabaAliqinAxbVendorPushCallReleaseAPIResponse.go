package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商推送通话结束事件 API返回值 
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel
}

// 供应商推送通话结束事件 成功返回结果
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_push_call_release_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaAliqinAxbVendorPushCallReleaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

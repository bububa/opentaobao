package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
转呼控制接口 APIResponse
alibaba.aliqin.axb.vendor.call.control

转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
*/
type AlibabaAliqinAxbVendorCallControlAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinAxbVendorCallControlResponse
}

type AlibabaAliqinAxbVendorCallControlResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_call_control_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 转呼控制接口响应
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}

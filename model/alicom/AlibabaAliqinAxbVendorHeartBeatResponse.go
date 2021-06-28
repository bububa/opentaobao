package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商心跳上报接口 APIResponse
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息
*/
type AlibabaAliqinAxbVendorHeartBeatAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinAxbVendorHeartBeatResponse
}

type AlibabaAliqinAxbVendorHeartBeatResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_heart_beat_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}

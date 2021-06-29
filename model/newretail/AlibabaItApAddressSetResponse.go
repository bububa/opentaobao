package newretail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
setApAddressNew APIResponse
alibaba.it.ap.address.set

该接口可为ISV系统提供 ap位置信息维护的功能
*/
type AlibabaItApAddressSetAPIResponse struct {
    model.CommonResponse
    AlibabaItApAddressSetResponse
}

type AlibabaItApAddressSetResponse struct {
    XMLName xml.Name `xml:"alibaba_it_ap_address_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaItApAddressSetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

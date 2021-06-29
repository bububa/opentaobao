package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV绑定外部门店id和外部商品id APIResponse
alibaba.alihealth.dental.item.bind

ISV绑定外部门店id和外部商品id
*/
type AlibabaAlihealthDentalItemBindAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalItemBindResponse
}

type AlibabaAlihealthDentalItemBindResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_item_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDentalItemBindMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

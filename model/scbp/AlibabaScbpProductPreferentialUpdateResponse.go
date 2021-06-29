package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设置P4P产品优先推广状态 APIResponse
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态
*/
type AlibabaScbpProductPreferentialUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpProductPreferentialUpdateResponse
}

type AlibabaScbpProductPreferentialUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_product_preferential_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设置成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

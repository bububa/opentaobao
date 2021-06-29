package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
替换橱窗商品 APIResponse
alibaba.scbp.showcase.updateproduct

替换橱窗商品
*/
type AlibabaScbpShowcaseUpdateproductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseUpdateproductResponse
}

type AlibabaScbpShowcaseUpdateproductResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_updateproduct_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否修改成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

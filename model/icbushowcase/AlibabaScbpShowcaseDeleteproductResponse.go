package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除橱窗商品 APIResponse
alibaba.scbp.showcase.deleteproduct

批量删除橱窗商品
*/
type AlibabaScbpShowcaseDeleteproductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseDeleteproductResponse
}

type AlibabaScbpShowcaseDeleteproductResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_deleteproduct_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

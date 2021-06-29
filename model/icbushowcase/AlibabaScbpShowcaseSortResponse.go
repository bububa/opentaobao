package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗顺序变更 APIResponse
alibaba.scbp.showcase.sort

橱窗顺序变更
*/
type AlibabaScbpShowcaseSortAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseSortResponse
}

type AlibabaScbpShowcaseSortResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_sort_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否更新成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}

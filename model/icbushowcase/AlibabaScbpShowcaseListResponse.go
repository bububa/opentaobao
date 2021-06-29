package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗查询 APIResponse
alibaba.scbp.showcase.list

橱窗查询
*/
type AlibabaScbpShowcaseListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseListResponse
}

type AlibabaScbpShowcaseListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Results   []Showcase `json:"results,omitempty" xml:"results>showcase,omitempty"`
    
    
}

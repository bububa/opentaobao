package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗状态 APIResponse
alibaba.scbp.showcase.status

查询橱窗状态，如总数、可用数量
*/
type AlibabaScbpShowcaseStatusAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseStatusResponse
}

type AlibabaScbpShowcaseStatusResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_status_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 全部橱窗数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 当前已用的橱窗数
    
    CurrentCount   int64 `json:"current_count,omitempty" xml:"current_count,omitempty"`

    
}

package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询对账单 APIResponse
alibaba.alihealth.dental.statement.query

ISV查询对账单
*/
type AlibabaAlihealthDentalStatementQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStatementQueryResponse
}

type AlibabaAlihealthDentalStatementQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_statement_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDentalStatementQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

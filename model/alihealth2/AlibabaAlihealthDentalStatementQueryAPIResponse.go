package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询对账单 API返回值 
alibaba.alihealth.dental.statement.query

ISV查询对账单
*/
type AlibabaAlihealthDentalStatementQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStatementQueryAPIResponseModel
}

// ISV查询对账单 成功返回结果
type AlibabaAlihealthDentalStatementQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_statement_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaAlihealthDentalStatementQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

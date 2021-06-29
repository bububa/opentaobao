package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询绑定审核状态 APIResponse
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态
*/
type AlibabaAlihealthDentalBindAuditQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalBindAuditQueryResponse
}

type AlibabaAlihealthDentalBindAuditQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_bind_audit_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDentalBindAuditQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

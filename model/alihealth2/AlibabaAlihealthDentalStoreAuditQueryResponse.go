package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店审核状态 APIResponse
alibaba.alihealth.dental.store.audit.query

ISV查询门店审核状态
*/
type AlibabaAlihealthDentalStoreAuditQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStoreAuditQueryResponse
}

type AlibabaAlihealthDentalStoreAuditQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_store_audit_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDentalStoreAuditQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

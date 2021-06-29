package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
异步开方处方查询 APIResponse
alibaba.alihealth.asyncprescribe.prescription.search

异步开方处方查询
*/
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAsyncprescribePrescriptionSearchResponse
}

type AlibabaAlihealthAsyncprescribePrescriptionSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_asyncprescribe_prescription_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}

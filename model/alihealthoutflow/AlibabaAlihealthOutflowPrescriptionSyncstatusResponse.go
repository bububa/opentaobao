package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-处方状态同步 APIResponse
alibaba.alihealth.outflow.prescription.syncstatus

阿里健康-处方外流-对外提供同步处方状态功能
*/
type AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowPrescriptionSyncstatusResponse
}

type AlibabaAlihealthOutflowPrescriptionSyncstatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_syncstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ServiceResult
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}

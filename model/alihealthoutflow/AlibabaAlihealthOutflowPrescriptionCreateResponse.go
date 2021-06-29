package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-创建处方 APIResponse
alibaba.alihealth.outflow.prescription.create

阿里健康-处方外流-对外提供保存处方功能
*/
type AlibabaAlihealthOutflowPrescriptionCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowPrescriptionCreateResponse
}

type AlibabaAlihealthOutflowPrescriptionCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ServiceResult
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}

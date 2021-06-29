package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方同步至医院返回校验结果 APIResponse
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果
*/
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowPrescriptionHospitalVerifyResponse
}

type AlibabaAlihealthOutflowPrescriptionHospitalVerifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_hospital_verify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ServiceResult
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}

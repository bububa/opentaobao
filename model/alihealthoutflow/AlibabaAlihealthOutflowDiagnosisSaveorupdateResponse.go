package alihealthoutflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-诊断字典表 APIResponse
alibaba.alihealth.outflow.diagnosis.saveorupdate

阿里健康-处方外流-对外提供诊断字典表维护功能
*/
type AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthOutflowDiagnosisSaveorupdateResponse
}

type AlibabaAlihealthOutflowDiagnosisSaveorupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_outflow_diagnosis_saveorupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ServiceResult
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}

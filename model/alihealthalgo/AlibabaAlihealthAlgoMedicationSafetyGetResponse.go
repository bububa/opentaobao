package alihealthalgo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药api APIResponse
alibaba.alihealth.algo.medication.safety.get

合理用药规则引擎服务
*/
type AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlgoMedicationSafetyGetResponse
}

type AlibabaAlihealthAlgoMedicationSafetyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_algo_medication_safety_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *SolutionResultTopSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}

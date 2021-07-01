package alihealthalgo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药api API返回值 
alibaba.alihealth.algo.medication.safety.get

合理用药规则引擎服务
*/
type AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel
}

// 合理用药api 成功返回结果
type AlibabaAlihealthAlgoMedicationSafetyGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_algo_medication_safety_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *SolutionResultTopSupport `json:"result,omitempty" xml:"result,omitempty"`
}

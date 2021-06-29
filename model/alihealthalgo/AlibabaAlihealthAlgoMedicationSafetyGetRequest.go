package alihealthalgo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药api APIRequest
alibaba.alihealth.algo.medication.safety.get

合理用药规则引擎服务
*/
type AlibabaAlihealthAlgoMedicationSafetyGetRequest struct {
    model.Params

    // 业务请求对象
    paramSolutionRequestTopSupport   *SolutionRequestTopSupport 

}

func NewAlibabaAlihealthAlgoMedicationSafetyGetRequest() *AlibabaAlihealthAlgoMedicationSafetyGetRequest{
    return &AlibabaAlihealthAlgoMedicationSafetyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.algo.medication.safety.get"
}

func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlgoMedicationSafetyGetRequest) SetParamSolutionRequestTopSupport(paramSolutionRequestTopSupport *SolutionRequestTopSupport) error {
    r.paramSolutionRequestTopSupport = paramSolutionRequestTopSupport
    r.Set("param_solution_request_top_support", paramSolutionRequestTopSupport)
    return nil
}

func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetParamSolutionRequestTopSupport() *SolutionRequestTopSupport {
    return r.paramSolutionRequestTopSupport
}


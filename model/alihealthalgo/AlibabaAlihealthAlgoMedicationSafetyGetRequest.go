package alihealthalgo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药api API请求
alibaba.alihealth.algo.medication.safety.get

合理用药规则引擎服务
*/
type AlibabaAlihealthAlgoMedicationSafetyGetRequest struct {
    model.Params
    // 业务请求对象
    _paramSolutionRequestTopSupport   *SolutionRequestTopSupport
}

// 初始化AlibabaAlihealthAlgoMedicationSafetyGetRequest对象
func NewAlibabaAlihealthAlgoMedicationSafetyGetRequest() *AlibabaAlihealthAlgoMedicationSafetyGetRequest{
    return &AlibabaAlihealthAlgoMedicationSafetyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.algo.medication.safety.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSolutionRequestTopSupport Setter
// 业务请求对象
func (r *AlibabaAlihealthAlgoMedicationSafetyGetRequest) SetParamSolutionRequestTopSupport(_paramSolutionRequestTopSupport *SolutionRequestTopSupport) error {
    r._paramSolutionRequestTopSupport = _paramSolutionRequestTopSupport
    r.Set("param_solution_request_top_support", _paramSolutionRequestTopSupport)
    return nil
}

// ParamSolutionRequestTopSupport Getter
func (r AlibabaAlihealthAlgoMedicationSafetyGetRequest) GetParamSolutionRequestTopSupport() *SolutionRequestTopSupport {
    return r._paramSolutionRequestTopSupport
}

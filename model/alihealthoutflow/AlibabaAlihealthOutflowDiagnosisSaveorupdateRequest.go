package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-诊断字典表 APIRequest
alibaba.alihealth.outflow.diagnosis.saveorupdate

阿里健康-处方外流-对外提供诊断字典表维护功能
*/
type AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest struct {
    model.Params

    // 诊断数据
    diagnoseDictDto   *DiagnoseDictDto 

}

func NewAlibabaAlihealthOutflowDiagnosisSaveorupdateRequest() *AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest{
    return &AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.diagnosis.saveorupdate"
}

func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) SetDiagnoseDictDto(diagnoseDictDto *DiagnoseDictDto) error {
    r.diagnoseDictDto = diagnoseDictDto
    r.Set("diagnose_dict_dto", diagnoseDictDto)
    return nil
}

func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetDiagnoseDictDto() *DiagnoseDictDto {
    return r.diagnoseDictDto
}


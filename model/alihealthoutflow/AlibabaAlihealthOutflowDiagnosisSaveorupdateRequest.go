package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-诊断字典表 API请求
alibaba.alihealth.outflow.diagnosis.saveorupdate

阿里健康-处方外流-对外提供诊断字典表维护功能
*/
type AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest struct {
    model.Params
    // 诊断数据
    _diagnoseDictDto   *DiagnoseDictDto
}

// 初始化AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest对象
func NewAlibabaAlihealthOutflowDiagnosisSaveorupdateRequest() *AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest{
    return &AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.diagnosis.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DiagnoseDictDto Setter
// 诊断数据
func (r *AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) SetDiagnoseDictDto(_diagnoseDictDto *DiagnoseDictDto) error {
    r._diagnoseDictDto = _diagnoseDictDto
    r.Set("diagnose_dict_dto", _diagnoseDictDto)
    return nil
}

// DiagnoseDictDto Getter
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateRequest) GetDiagnoseDictDto() *DiagnoseDictDto {
    return r._diagnoseDictDto
}

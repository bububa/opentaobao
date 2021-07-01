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
type AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest struct {
    model.Params
    // 诊断数据
    _diagnoseDictDto   *DiagnoseDictDto
}

// 初始化AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest对象
func NewAlibabaAlihealthOutflowDiagnosisSaveorupdateRequest() *AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest{
    return &AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.diagnosis.saveorupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DiagnoseDictDto Setter
// 诊断数据
func (r *AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest) SetDiagnoseDictDto(_diagnoseDictDto *DiagnoseDictDto) error {
    r._diagnoseDictDto = _diagnoseDictDto
    r.Set("diagnose_dict_dto", _diagnoseDictDto)
    return nil
}

// DiagnoseDictDto Getter
func (r AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest) GetDiagnoseDictDto() *DiagnoseDictDto {
    return r._diagnoseDictDto
}

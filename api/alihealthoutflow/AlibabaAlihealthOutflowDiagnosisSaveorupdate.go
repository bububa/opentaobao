package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流-诊断字典表 
alibaba.alihealth.outflow.diagnosis.saveorupdate

阿里健康-处方外流-对外提供诊断字典表维护功能
*/
func AlibabaAlihealthOutflowDiagnosisSaveorupdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowDiagnosisSaveorupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

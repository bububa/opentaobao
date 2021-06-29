package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流-用法字典表 
alibaba.alihealth.outflow.usage.saveorupdate

阿里健康-处方外流-对外提供用法字典表维护功能
*/
func AlibabaAlihealthOutflowUsageSaveorupdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowUsageSaveorupdateRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowUsageSaveorupdateAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowUsageSaveorupdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

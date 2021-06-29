package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流-处方状态同步 
alibaba.alihealth.outflow.prescription.syncstatus

阿里健康-处方外流-对外提供同步处方状态功能
*/
func AlibabaAlihealthOutflowPrescriptionSyncstatus(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionSyncstatusRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowPrescriptionSyncstatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
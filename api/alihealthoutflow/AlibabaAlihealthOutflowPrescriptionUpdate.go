package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方外流-修改处方 
alibaba.alihealth.outflow.prescription.update

阿里健康-处方外流-对外提供处方修改功能
*/
func AlibabaAlihealthOutflowPrescriptionUpdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionUpdateRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowPrescriptionUpdateAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowPrescriptionUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

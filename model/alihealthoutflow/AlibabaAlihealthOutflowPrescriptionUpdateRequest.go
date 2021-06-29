package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-修改处方 APIRequest
alibaba.alihealth.outflow.prescription.update

阿里健康-处方外流-对外提供处方修改功能
*/
type AlibabaAlihealthOutflowPrescriptionUpdateRequest struct {
    model.Params

    // 入参对象
    updateRequest   *PrescriptionOutflowUpdateRequest 

}

func NewAlibabaAlihealthOutflowPrescriptionUpdateRequest() *AlibabaAlihealthOutflowPrescriptionUpdateRequest{
    return &AlibabaAlihealthOutflowPrescriptionUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.update"
}

func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowPrescriptionUpdateRequest) SetUpdateRequest(updateRequest *PrescriptionOutflowUpdateRequest) error {
    r.updateRequest = updateRequest
    r.Set("update_request", updateRequest)
    return nil
}

func (r AlibabaAlihealthOutflowPrescriptionUpdateRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
    return r.updateRequest
}


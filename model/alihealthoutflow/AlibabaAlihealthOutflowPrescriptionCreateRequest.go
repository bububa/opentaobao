package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流-创建处方 APIRequest
alibaba.alihealth.outflow.prescription.create

阿里健康-处方外流-对外提供保存处方功能
*/
type AlibabaAlihealthOutflowPrescriptionCreateRequest struct {
    model.Params

    // 保存处方入参
    createRequest   *PrescriptionOutflowUpdateRequest 

}

func NewAlibabaAlihealthOutflowPrescriptionCreateRequest() *AlibabaAlihealthOutflowPrescriptionCreateRequest{
    return &AlibabaAlihealthOutflowPrescriptionCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.create"
}

func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowPrescriptionCreateRequest) SetCreateRequest(createRequest *PrescriptionOutflowUpdateRequest) error {
    r.createRequest = createRequest
    r.Set("create_request", createRequest)
    return nil
}

func (r AlibabaAlihealthOutflowPrescriptionCreateRequest) GetCreateRequest() *PrescriptionOutflowUpdateRequest {
    return r.createRequest
}


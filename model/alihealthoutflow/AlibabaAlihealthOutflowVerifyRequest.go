package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流药店通过核销码核销处方 APIRequest
alibaba.alihealth.outflow.verify

处方外流药店通过核销码核销处方
*/
type AlibabaAlihealthOutflowVerifyRequest struct {
    model.Params

    // 入参
    prescriptionVerifyRequest   *PrescriptionVerifyRequest 

}

func NewAlibabaAlihealthOutflowVerifyRequest() *AlibabaAlihealthOutflowVerifyRequest{
    return &AlibabaAlihealthOutflowVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.verify"
}

func (r AlibabaAlihealthOutflowVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowVerifyRequest) SetPrescriptionVerifyRequest(prescriptionVerifyRequest *PrescriptionVerifyRequest) error {
    r.prescriptionVerifyRequest = prescriptionVerifyRequest
    r.Set("prescription_verify_request", prescriptionVerifyRequest)
    return nil
}

func (r AlibabaAlihealthOutflowVerifyRequest) GetPrescriptionVerifyRequest() *PrescriptionVerifyRequest {
    return r.prescriptionVerifyRequest
}


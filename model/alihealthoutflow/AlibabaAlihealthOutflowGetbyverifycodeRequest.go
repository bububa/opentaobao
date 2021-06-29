package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流药店通过核销码获取处方 APIRequest
alibaba.alihealth.outflow.getbyverifycode

阿里健康对合作药店提供通过核销码查看处方的功能
*/
type AlibabaAlihealthOutflowGetbyverifycodeRequest struct {
    model.Params

    // 入参
    prescriptionGetByVerifyRequest   *PrescriptionGetByVerifyRequest 

}

func NewAlibabaAlihealthOutflowGetbyverifycodeRequest() *AlibabaAlihealthOutflowGetbyverifycodeRequest{
    return &AlibabaAlihealthOutflowGetbyverifycodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowGetbyverifycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.getbyverifycode"
}

func (r AlibabaAlihealthOutflowGetbyverifycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowGetbyverifycodeRequest) SetPrescriptionGetByVerifyRequest(prescriptionGetByVerifyRequest *PrescriptionGetByVerifyRequest) error {
    r.prescriptionGetByVerifyRequest = prescriptionGetByVerifyRequest
    r.Set("prescription_get_by_verify_request", prescriptionGetByVerifyRequest)
    return nil
}

func (r AlibabaAlihealthOutflowGetbyverifycodeRequest) GetPrescriptionGetByVerifyRequest() *PrescriptionGetByVerifyRequest {
    return r.prescriptionGetByVerifyRequest
}


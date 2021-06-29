package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方同步至医院返回校验结果 APIRequest
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果
*/
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest struct {
    model.Params

    // 入参对象
    updateRequest   *PrescriptionOutflowUpdateRequest 

}

func NewAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest() *AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest{
    return &AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.hospital.verify"
}

func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) SetUpdateRequest(updateRequest *PrescriptionOutflowUpdateRequest) error {
    r.updateRequest = updateRequest
    r.Set("update_request", updateRequest)
    return nil
}

func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
    return r.updateRequest
}


package alihealth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康处方平台获取授权码 APIRequest
alibaba.alihealth.prescription.auth.get

获取处方用户授权
*/
type AlibabaAlihealthPrescriptionAuthGetRequest struct {
    model.Params

    // 请求入参
    prescriptionRequest   *PrescriptionDoctorAuthRequest 

}

func NewAlibabaAlihealthPrescriptionAuthGetRequest() *AlibabaAlihealthPrescriptionAuthGetRequest{
    return &AlibabaAlihealthPrescriptionAuthGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.prescription.auth.get"
}

func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPrescriptionAuthGetRequest) SetPrescriptionRequest(prescriptionRequest *PrescriptionDoctorAuthRequest) error {
    r.prescriptionRequest = prescriptionRequest
    r.Set("prescription_request", prescriptionRequest)
    return nil
}

func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetPrescriptionRequest() *PrescriptionDoctorAuthRequest {
    return r.prescriptionRequest
}


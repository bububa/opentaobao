package alihealth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康处方平台获取授权码 API请求
alibaba.alihealth.prescription.auth.get

获取处方用户授权
*/
type AlibabaAlihealthPrescriptionAuthGetRequest struct {
    model.Params
    // 请求入参
    _prescriptionRequest   *PrescriptionDoctorAuthRequest
}

// 初始化AlibabaAlihealthPrescriptionAuthGetRequest对象
func NewAlibabaAlihealthPrescriptionAuthGetRequest() *AlibabaAlihealthPrescriptionAuthGetRequest{
    return &AlibabaAlihealthPrescriptionAuthGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.prescription.auth.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PrescriptionRequest Setter
// 请求入参
func (r *AlibabaAlihealthPrescriptionAuthGetRequest) SetPrescriptionRequest(_prescriptionRequest *PrescriptionDoctorAuthRequest) error {
    r._prescriptionRequest = _prescriptionRequest
    r.Set("prescription_request", _prescriptionRequest)
    return nil
}

// PrescriptionRequest Getter
func (r AlibabaAlihealthPrescriptionAuthGetRequest) GetPrescriptionRequest() *PrescriptionDoctorAuthRequest {
    return r._prescriptionRequest
}

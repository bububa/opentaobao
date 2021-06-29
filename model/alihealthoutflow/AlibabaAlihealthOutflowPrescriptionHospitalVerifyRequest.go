package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方同步至医院返回校验结果 API请求
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果
*/
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest struct {
    model.Params
    // 入参对象
    _updateRequest   *PrescriptionOutflowUpdateRequest
}

// 初始化AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest对象
func NewAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest() *AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest{
    return &AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.prescription.hospital.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateRequest Setter
// 入参对象
func (r *AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
    r._updateRequest = _updateRequest
    r.Set("update_request", _updateRequest)
    return nil
}

// UpdateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
    return r._updateRequest
}

package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流药店通过核销码获取处方 API请求
alibaba.alihealth.outflow.getbyverifycode

阿里健康对合作药店提供通过核销码查看处方的功能
*/
type AlibabaAlihealthOutflowGetbyverifycodeAPIRequest struct {
    model.Params
    // 入参
    _prescriptionGetByVerifyRequest   *PrescriptionGetByVerifyRequest
}

// 初始化AlibabaAlihealthOutflowGetbyverifycodeAPIRequest对象
func NewAlibabaAlihealthOutflowGetbyverifycodeRequest() *AlibabaAlihealthOutflowGetbyverifycodeAPIRequest{
    return &AlibabaAlihealthOutflowGetbyverifycodeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowGetbyverifycodeAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.getbyverifycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowGetbyverifycodeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PrescriptionGetByVerifyRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowGetbyverifycodeAPIRequest) SetPrescriptionGetByVerifyRequest(_prescriptionGetByVerifyRequest *PrescriptionGetByVerifyRequest) error {
    r._prescriptionGetByVerifyRequest = _prescriptionGetByVerifyRequest
    r.Set("prescription_get_by_verify_request", _prescriptionGetByVerifyRequest)
    return nil
}

// PrescriptionGetByVerifyRequest Getter
func (r AlibabaAlihealthOutflowGetbyverifycodeAPIRequest) GetPrescriptionGetByVerifyRequest() *PrescriptionGetByVerifyRequest {
    return r._prescriptionGetByVerifyRequest
}

package alihealthoutflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处方外流药店通过核销码核销处方 API请求
alibaba.alihealth.outflow.verify

处方外流药店通过核销码核销处方
*/
type AlibabaAlihealthOutflowVerifyAPIRequest struct {
    model.Params
    // 入参
    _prescriptionVerifyRequest   *PrescriptionVerifyRequest
}

// 初始化AlibabaAlihealthOutflowVerifyAPIRequest对象
func NewAlibabaAlihealthOutflowVerifyRequest() *AlibabaAlihealthOutflowVerifyAPIRequest{
    return &AlibabaAlihealthOutflowVerifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PrescriptionVerifyRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVerifyAPIRequest) SetPrescriptionVerifyRequest(_prescriptionVerifyRequest *PrescriptionVerifyRequest) error {
    r._prescriptionVerifyRequest = _prescriptionVerifyRequest
    r.Set("prescription_verify_request", _prescriptionVerifyRequest)
    return nil
}

// PrescriptionVerifyRequest Getter
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetPrescriptionVerifyRequest() *PrescriptionVerifyRequest {
    return r._prescriptionVerifyRequest
}

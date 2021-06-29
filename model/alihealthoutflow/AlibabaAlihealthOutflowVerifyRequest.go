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
type AlibabaAlihealthOutflowVerifyRequest struct {
    model.Params
    // 入参
    prescriptionVerifyRequest   *PrescriptionVerifyRequest
}

// 初始化AlibabaAlihealthOutflowVerifyRequest对象
func NewAlibabaAlihealthOutflowVerifyRequest() *AlibabaAlihealthOutflowVerifyRequest{
    return &AlibabaAlihealthOutflowVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.outflow.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PrescriptionVerifyRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVerifyRequest) SetPrescriptionVerifyRequest(prescriptionVerifyRequest *PrescriptionVerifyRequest) error {
    r.prescriptionVerifyRequest = prescriptionVerifyRequest
    r.Set("prescription_verify_request", prescriptionVerifyRequest)
    return nil
}

// PrescriptionVerifyRequest Getter
func (r AlibabaAlihealthOutflowVerifyRequest) GetPrescriptionVerifyRequest() *PrescriptionVerifyRequest {
    return r.prescriptionVerifyRequest
}

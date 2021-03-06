package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowVerifyAPIRequest 处方外流药店通过核销码核销处方 API请求
// alibaba.alihealth.outflow.verify
//
// 处方外流药店通过核销码核销处方
type AlibabaAlihealthOutflowVerifyAPIRequest struct {
	model.Params
	// 入参
	_prescriptionVerifyRequest *PrescriptionVerifyRequest
}

// NewAlibabaAlihealthOutflowVerifyRequest 初始化AlibabaAlihealthOutflowVerifyAPIRequest对象
func NewAlibabaAlihealthOutflowVerifyRequest() *AlibabaAlihealthOutflowVerifyAPIRequest {
	return &AlibabaAlihealthOutflowVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPrescriptionVerifyRequest is PrescriptionVerifyRequest Setter
// 入参
func (r *AlibabaAlihealthOutflowVerifyAPIRequest) SetPrescriptionVerifyRequest(_prescriptionVerifyRequest *PrescriptionVerifyRequest) error {
	r._prescriptionVerifyRequest = _prescriptionVerifyRequest
	r.Set("prescription_verify_request", _prescriptionVerifyRequest)
	return nil
}

// GetPrescriptionVerifyRequest PrescriptionVerifyRequest Getter
func (r AlibabaAlihealthOutflowVerifyAPIRequest) GetPrescriptionVerifyRequest() *PrescriptionVerifyRequest {
	return r._prescriptionVerifyRequest
}

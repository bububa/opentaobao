package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowVerifyAPIRequest
处方外流药店通过核销码核销处方 API请求
alibaba.alihealth.outflow.verify

处方外流药店通过核销码核销处方 */
type AlibabaAlihealthOutflowVerifyAPIRequest struct {
	model.Params
	// 入参
	_prescriptionVerifyRequest *PrescriptionVerifyRequest
}

// New

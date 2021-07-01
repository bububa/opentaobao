package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest
处方同步至医院返回校验结果 API请求
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果 */
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// New

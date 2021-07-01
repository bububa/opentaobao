package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPrescriptionCreateAPIRequest
处方外流-创建处方 API请求
alibaba.alihealth.outflow.prescription.create

阿里健康-处方外流-对外提供保存处方功能 */
type AlibabaAlihealthOutflowPrescriptionCreateAPIRequest struct {
	model.Params
	// 保存处方入参
	_createRequest *PrescriptionOutflowUpdateRequest
}

// New

package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest
处方外流-修改处方 API请求
alibaba.alihealth.outflow.prescription.update

阿里健康-处方外流-对外提供处方修改功能 */
type AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// New

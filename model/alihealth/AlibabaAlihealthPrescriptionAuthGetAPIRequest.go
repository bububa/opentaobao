package alihealth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPrescriptionAuthGetAPIRequest
阿里健康处方平台获取授权码 API请求
alibaba.alihealth.prescription.auth.get

获取处方用户授权 */
type AlibabaAlihealthPrescriptionAuthGetAPIRequest struct {
	model.Params
	// 请求入参
	_prescriptionRequest *PrescriptionDoctorAuthRequest
}

// New

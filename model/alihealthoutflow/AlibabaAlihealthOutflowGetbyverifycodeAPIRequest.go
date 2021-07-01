package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowGetbyverifycodeAPIRequest
处方外流药店通过核销码获取处方 API请求
alibaba.alihealth.outflow.getbyverifycode

阿里健康对合作药店提供通过核销码查看处方的功能 */
type AlibabaAlihealthOutflowGetbyverifycodeAPIRequest struct {
	model.Params
	// 入参
	_prescriptionGetByVerifyRequest *PrescriptionGetByVerifyRequest
}

// New

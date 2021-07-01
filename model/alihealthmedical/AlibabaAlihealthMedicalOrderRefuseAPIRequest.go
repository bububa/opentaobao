package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalOrderRefuseAPIRequest
三方机构通知平台"医生拒诊" API请求
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊" */
type AlibabaAlihealthMedicalOrderRefuseAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *RefuseOrderRequestDto
}

// New

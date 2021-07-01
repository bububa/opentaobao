package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalRegistrationSyncAPIRequest
阿里健康支付宝挂号记录回传接口 API请求
alibaba.alihealth.medical.registration.sync

阿里健康支付宝挂号记录回传接口 */
type AlibabaAlihealthMedicalRegistrationSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// New

package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDoctorSyncAPIRequest
阿里健康预约挂号医生同步接口 API请求
alibaba.alihealth.medical.doctor.sync

阿里健康预约挂号医生同步接口 */
type AlibabaAlihealthMedicalDoctorSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// New

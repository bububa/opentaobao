package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest
阿里健康新挂号数据回传 API请求
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口 */
type AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// New

package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDepartmentSyncAPIRequest
阿里健康预约挂号科室同步接口 API请求
alibaba.alihealth.medical.department.sync

阿里健康预约挂号科室同步接口 */
type AlibabaAlihealthMedicalDepartmentSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// New

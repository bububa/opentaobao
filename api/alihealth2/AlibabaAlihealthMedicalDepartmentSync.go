package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalDepartmentSync 阿里健康预约挂号科室同步接口
// alibaba.alihealth.medical.department.sync
//
// 阿里健康预约挂号科室同步接口
func AlibabaAlihealthMedicalDepartmentSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalDepartmentSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalDepartmentSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

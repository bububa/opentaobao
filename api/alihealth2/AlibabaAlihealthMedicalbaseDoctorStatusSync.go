package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDoctorStatusSync 挂号医生上下架
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseDoctorStatusSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalDoctorSync 阿里健康预约挂号医生同步接口
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
func AlibabaAlihealthMedicalDoctorSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalDoctorSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalDoctorSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

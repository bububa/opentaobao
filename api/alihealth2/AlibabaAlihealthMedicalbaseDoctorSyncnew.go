package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDoctorSyncnew 直连医生上传
// alibaba.alihealth.medicalbase.doctor.syncnew
//
// 直连医生上传
func AlibabaAlihealthMedicalbaseDoctorSyncnew(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDoctorSyncnewAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

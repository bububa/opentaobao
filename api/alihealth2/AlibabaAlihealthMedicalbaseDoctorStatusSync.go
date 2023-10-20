package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDoctorStatusSync 挂号医生上下架
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseDoctorStatusSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

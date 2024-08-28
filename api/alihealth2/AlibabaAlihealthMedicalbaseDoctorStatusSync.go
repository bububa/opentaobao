package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDoctorStatusSync 挂号医生上下架
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseDoctorStatusSync(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

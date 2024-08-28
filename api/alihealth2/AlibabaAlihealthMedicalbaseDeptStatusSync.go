package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseDeptStatusSync 挂号科室上下线
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseDeptStatusSync(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseDeptStatusSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseDeptStatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

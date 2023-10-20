package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseHosStatusSync 挂号医院上下线
// alibaba.alihealth.medicalbase.hos.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseHosStatusSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

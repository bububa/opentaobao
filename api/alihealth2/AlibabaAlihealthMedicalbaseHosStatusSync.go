package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseHosStatusSync 挂号医院上下线
// alibaba.alihealth.medicalbase.hos.status.sync
//
// 挂号医院上下线
func AlibabaAlihealthMedicalbaseHosStatusSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalbaseHosStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

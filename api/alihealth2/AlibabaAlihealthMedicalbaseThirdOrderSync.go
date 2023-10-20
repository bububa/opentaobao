package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseThirdOrderSync 第三方订单同步
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
func AlibabaAlihealthMedicalbaseThirdOrderSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

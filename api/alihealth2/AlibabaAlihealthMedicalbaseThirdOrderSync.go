package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalbaseThirdOrderSync 第三方订单同步
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
func AlibabaAlihealthMedicalbaseThirdOrderSync(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

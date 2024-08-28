package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalRegistrationSync 阿里健康支付宝挂号记录回传接口
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
func AlibabaAlihealthMedicalRegistrationSync(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalRegistrationSyncAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalRegistrationSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

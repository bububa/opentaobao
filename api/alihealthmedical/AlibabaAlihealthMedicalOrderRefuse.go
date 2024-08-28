package alihealthmedical

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalOrderRefuse 三方机构通知平台"医生拒诊"
// alibaba.alihealth.medical.order.refuse
//
// 三方机构通知平台&#34;医生拒诊&#34;
func AlibabaAlihealthMedicalOrderRefuse(ctx context.Context, clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

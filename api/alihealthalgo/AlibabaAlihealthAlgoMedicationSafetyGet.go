package alihealthalgo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthalgo"
)

// AlibabaAlihealthAlgoMedicationSafetyGet 合理用药api
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
func AlibabaAlihealthAlgoMedicationSafetyGet(ctx context.Context, clt *core.SDKClient, req *alihealthalgo.AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest, resp *alihealthalgo.AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

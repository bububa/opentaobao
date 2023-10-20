package alihealthalgo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthalgo"
)

// AlibabaAlihealthAlgoMedicationSafetyGet 合理用药api
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
func AlibabaAlihealthAlgoMedicationSafetyGet(clt *core.SDKClient, req *alihealthalgo.AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest, resp *alihealthalgo.AlibabaAlihealthAlgoMedicationSafetyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

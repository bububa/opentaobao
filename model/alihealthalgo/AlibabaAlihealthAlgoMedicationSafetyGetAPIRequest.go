package alihealthalgo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest
合理用药api API请求
alibaba.alihealth.algo.medication.safety.get

合理用药规则引擎服务 */
type AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest struct {
	model.Params
	// 业务请求对象
	_paramSolutionRequestTopSupport *SolutionRequestTopSupport
}

// New

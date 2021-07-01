package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest
保险退货服务商勘察结论提交接口 API请求
alipay.baoxian.claim.survey.conclusion.submit

保险退货服务商提交勘察结论 */
type AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest struct {
	model.Params
	// 勘察结论
	_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto
}

// New

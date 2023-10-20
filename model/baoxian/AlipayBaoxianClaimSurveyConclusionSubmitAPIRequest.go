package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest 保险退货服务商勘察结论提交接口 API请求
// alipay.baoxian.claim.survey.conclusion.submit
//
// 保险退货服务商提交勘察结论
type AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest struct {
	model.Params
	// 勘察结论
	_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto
}

// NewAlipaybaoxianclaimsurveyconclusionsubmitRequest 初始化AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest对象
func NewAlipaybaoxianclaimsurveyconclusionsubmitRequest() *AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest {
	return &AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.survey.conclusion.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFacilitatorSurveyConclusions is FacilitatorSurveyConclusions Setter
// 勘察结论
func (r *AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest) SetFacilitatorSurveyConclusions(_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto) error {
	r._facilitatorSurveyConclusions = _facilitatorSurveyConclusions
	r.Set("facilitator_survey_conclusions", _facilitatorSurveyConclusions)
	return nil
}

// GetFacilitatorSurveyConclusions FacilitatorSurveyConclusions Getter
func (r AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest) GetFacilitatorSurveyConclusions() []InsFacilitatorSurveyConclusionDto {
	return r._facilitatorSurveyConclusions
}

package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保险退货服务商勘察结论提交接口 API请求
alipay.baoxian.claim.survey.conclusion.submit

保险退货服务商提交勘察结论
*/
type AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest struct {
    model.Params
    // 勘察结论
    _facilitatorSurveyConclusions   []InsFacilitatorSurveyConclusionDto
}

// 初始化AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest对象
func NewAlipayBaoxianClaimSurveyConclusionSubmitRequest() *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest{
    return &AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.survey.conclusion.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FacilitatorSurveyConclusions Setter
// 勘察结论
func (r *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) SetFacilitatorSurveyConclusions(_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto) error {
    r._facilitatorSurveyConclusions = _facilitatorSurveyConclusions
    r.Set("facilitator_survey_conclusions", _facilitatorSurveyConclusions)
    return nil
}

// FacilitatorSurveyConclusions Getter
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetFacilitatorSurveyConclusions() []InsFacilitatorSurveyConclusionDto {
    return r._facilitatorSurveyConclusions
}

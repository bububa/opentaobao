package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保险退货服务商勘察结论提交接口 APIRequest
alipay.baoxian.claim.survey.conclusion.submit

保险退货服务商提交勘察结论
*/
type AlipayBaoxianClaimSurveyConclusionSubmitRequest struct {
    model.Params

    // 勘察结论
    facilitatorSurveyConclusions   []InsFacilitatorSurveyConclusionDto 

}

func NewAlipayBaoxianClaimSurveyConclusionSubmitRequest() *AlipayBaoxianClaimSurveyConclusionSubmitRequest{
    return &AlipayBaoxianClaimSurveyConclusionSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlipayBaoxianClaimSurveyConclusionSubmitRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.survey.conclusion.submit"
}

func (r AlipayBaoxianClaimSurveyConclusionSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlipayBaoxianClaimSurveyConclusionSubmitRequest) SetFacilitatorSurveyConclusions(facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto) error {
    r.facilitatorSurveyConclusions = facilitatorSurveyConclusions
    r.Set("facilitator_survey_conclusions", facilitatorSurveyConclusions)
    return nil
}

func (r AlipayBaoxianClaimSurveyConclusionSubmitRequest) GetFacilitatorSurveyConclusions() []InsFacilitatorSurveyConclusionDto {
    return r.facilitatorSurveyConclusions
}


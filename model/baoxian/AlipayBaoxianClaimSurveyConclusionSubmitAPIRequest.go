package baoxian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest 保险退货服务商勘察结论提交接口 API请求
// alipay.baoxian.claim.survey.conclusion.submit
//
// 保险退货服务商提交勘察结论
type AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest struct {
	model.Params
	// 勘察结论
	_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto
}

// NewAlipayBaoxianClaimSurveyConclusionSubmitRequest 初始化AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest对象
func NewAlipayBaoxianClaimSurveyConclusionSubmitRequest() *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest {
	return &AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) Reset() {
	r._facilitatorSurveyConclusions = r._facilitatorSurveyConclusions[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.survey.conclusion.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFacilitatorSurveyConclusions is FacilitatorSurveyConclusions Setter
// 勘察结论
func (r *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) SetFacilitatorSurveyConclusions(_facilitatorSurveyConclusions []InsFacilitatorSurveyConclusionDto) error {
	r._facilitatorSurveyConclusions = _facilitatorSurveyConclusions
	r.Set("facilitator_survey_conclusions", _facilitatorSurveyConclusions)
	return nil
}

// GetFacilitatorSurveyConclusions FacilitatorSurveyConclusions Getter
func (r AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) GetFacilitatorSurveyConclusions() []InsFacilitatorSurveyConclusionDto {
	return r._facilitatorSurveyConclusions
}

var poolAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlipayBaoxianClaimSurveyConclusionSubmitRequest()
	},
}

// GetAlipayBaoxianClaimSurveyConclusionSubmitRequest 从 sync.Pool 获取 AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest
func GetAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest() *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest {
	return poolAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest.Get().(*AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest)
}

// ReleaseAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest 将 AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest 放入 sync.Pool
func ReleaseAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest(v *AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest) {
	v.Reset()
	poolAlipayBaoxianClaimSurveyConclusionSubmitAPIRequest.Put(v)
}

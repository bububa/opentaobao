package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest 三方评论信息同步 API请求
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
type AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest struct {
	model.Params
	// 请求参数
	_evaluateRequest *MedicalBaseTopRequestDto
}

// NewAlibabaalihealthmedicalbasethirdevaluatesyncRequest 初始化AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest对象
func NewAlibabaalihealthmedicalbasethirdevaluatesyncRequest() *AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest {
	return &AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.third.evaluate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEvaluateRequest is EvaluateRequest Setter
// 请求参数
func (r *AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest) SetEvaluateRequest(_evaluateRequest *MedicalBaseTopRequestDto) error {
	r._evaluateRequest = _evaluateRequest
	r.Set("evaluate_request", _evaluateRequest)
	return nil
}

// GetEvaluateRequest EvaluateRequest Getter
func (r AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest) GetEvaluateRequest() *MedicalBaseTopRequestDto {
	return r._evaluateRequest
}

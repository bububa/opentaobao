package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest 三方评论信息同步 API请求
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest struct {
	model.Params
	// 请求参数
	_evaluateRequest *MedicalBaseTopRequestDto
}

// NewAlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest 初始化AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest() *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.third.evaluate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEvaluateRequest is EvaluateRequest Setter
// 请求参数
func (r *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) SetEvaluateRequest(_evaluateRequest *MedicalBaseTopRequestDto) error {
	r._evaluateRequest = _evaluateRequest
	r.Set("evaluate_request", _evaluateRequest)
	return nil
}

// GetEvaluateRequest EvaluateRequest Getter
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetEvaluateRequest() *MedicalBaseTopRequestDto {
	return r._evaluateRequest
}

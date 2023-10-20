package alihealth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) Reset() {
	r._evaluateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.third.evaluate.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest() *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest 将 AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest(v *AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIRequest.Put(v)
}

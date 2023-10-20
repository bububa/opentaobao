package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest 15-供应商反馈（原料）同步接口 API请求
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest 初始化AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest() *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanFeedbackRawUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest
func GetAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest() *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest 将 AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest(v *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest.Put(v)
}

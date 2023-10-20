package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest 14-供应商反馈（OEM）同步接口 API请求
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
type AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest 初始化AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest() *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanFeedbackOemUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest
func GetAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest() *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest 将 AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest(v *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest.Put(v)
}

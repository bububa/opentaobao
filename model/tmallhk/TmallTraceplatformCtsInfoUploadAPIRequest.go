package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCtsInfoUploadAPIRequest CTS提交溯源信息 API请求
// tmall.traceplatform.cts.info.upload
//
// cts上传溯源信息
type TmallTraceplatformCtsInfoUploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *CtsInfo
}

// NewTmallTraceplatformCtsInfoUploadRequest 初始化TmallTraceplatformCtsInfoUploadAPIRequest对象
func NewTmallTraceplatformCtsInfoUploadRequest() *TmallTraceplatformCtsInfoUploadAPIRequest {
	return &TmallTraceplatformCtsInfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsInfoUploadAPIRequest) SetTraceInfo(_traceInfo *CtsInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// Get TraceInfo Getter
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetTraceInfo() *CtsInfo {
	return r._traceInfo
}

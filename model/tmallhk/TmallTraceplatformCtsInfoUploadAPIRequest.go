package tmallhk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraceplatformCtsInfoUploadAPIRequest) Reset() {
	r._traceInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsInfoUploadAPIRequest) SetTraceInfo(_traceInfo *CtsInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmallTraceplatformCtsInfoUploadAPIRequest) GetTraceInfo() *CtsInfo {
	return r._traceInfo
}

var poolTmallTraceplatformCtsInfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraceplatformCtsInfoUploadRequest()
	},
}

// GetTmallTraceplatformCtsInfoUploadRequest 从 sync.Pool 获取 TmallTraceplatformCtsInfoUploadAPIRequest
func GetTmallTraceplatformCtsInfoUploadAPIRequest() *TmallTraceplatformCtsInfoUploadAPIRequest {
	return poolTmallTraceplatformCtsInfoUploadAPIRequest.Get().(*TmallTraceplatformCtsInfoUploadAPIRequest)
}

// ReleaseTmallTraceplatformCtsInfoUploadAPIRequest 将 TmallTraceplatformCtsInfoUploadAPIRequest 放入 sync.Pool
func ReleaseTmallTraceplatformCtsInfoUploadAPIRequest(v *TmallTraceplatformCtsInfoUploadAPIRequest) {
	v.Reset()
	poolTmallTraceplatformCtsInfoUploadAPIRequest.Put(v)
}

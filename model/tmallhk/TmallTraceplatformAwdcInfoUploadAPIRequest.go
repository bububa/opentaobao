package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformAwdcInfoUploadAPIRequest AWDC提交溯源信息 API请求
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
type TmallTraceplatformAwdcInfoUploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *AwdcInfo
}

// NewTmallTraceplatformAwdcInfoUploadRequest 初始化TmallTraceplatformAwdcInfoUploadAPIRequest对象
func NewTmallTraceplatformAwdcInfoUploadRequest() *TmallTraceplatformAwdcInfoUploadAPIRequest {
	return &TmallTraceplatformAwdcInfoUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraceplatformAwdcInfoUploadAPIRequest) Reset() {
	r._traceInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.awdc.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformAwdcInfoUploadAPIRequest) SetTraceInfo(_traceInfo *AwdcInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetTraceInfo() *AwdcInfo {
	return r._traceInfo
}

var poolTmallTraceplatformAwdcInfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraceplatformAwdcInfoUploadRequest()
	},
}

// GetTmallTraceplatformAwdcInfoUploadRequest 从 sync.Pool 获取 TmallTraceplatformAwdcInfoUploadAPIRequest
func GetTmallTraceplatformAwdcInfoUploadAPIRequest() *TmallTraceplatformAwdcInfoUploadAPIRequest {
	return poolTmallTraceplatformAwdcInfoUploadAPIRequest.Get().(*TmallTraceplatformAwdcInfoUploadAPIRequest)
}

// ReleaseTmallTraceplatformAwdcInfoUploadAPIRequest 将 TmallTraceplatformAwdcInfoUploadAPIRequest 放入 sync.Pool
func ReleaseTmallTraceplatformAwdcInfoUploadAPIRequest(v *TmallTraceplatformAwdcInfoUploadAPIRequest) {
	v.Reset()
	poolTmallTraceplatformAwdcInfoUploadAPIRequest.Put(v)
}

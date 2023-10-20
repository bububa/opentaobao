package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformctsinfouploadAPIRequest CTS提交溯源信息 API请求
// tmall.traceplatform.cts.info.upload
//
// cts上传溯源信息
type TmalltraceplatformctsinfouploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *CtsInfo
}

// NewTmalltraceplatformctsinfouploadRequest 初始化TmalltraceplatformctsinfouploadAPIRequest对象
func NewTmalltraceplatformctsinfouploadRequest() *TmalltraceplatformctsinfouploadAPIRequest {
	return &TmalltraceplatformctsinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraceplatformctsinfouploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraceplatformctsinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraceplatformctsinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmalltraceplatformctsinfouploadAPIRequest) SetTraceInfo(_traceInfo *CtsInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmalltraceplatformctsinfouploadAPIRequest) GetTraceInfo() *CtsInfo {
	return r._traceInfo
}

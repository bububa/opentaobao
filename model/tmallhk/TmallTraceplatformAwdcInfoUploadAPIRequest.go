package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformawdcinfouploadAPIRequest AWDC提交溯源信息 API请求
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
type TmalltraceplatformawdcinfouploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *AwdcInfo
}

// NewTmalltraceplatformawdcinfouploadRequest 初始化TmalltraceplatformawdcinfouploadAPIRequest对象
func NewTmalltraceplatformawdcinfouploadRequest() *TmalltraceplatformawdcinfouploadAPIRequest {
	return &TmalltraceplatformawdcinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraceplatformawdcinfouploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.awdc.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraceplatformawdcinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraceplatformawdcinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmalltraceplatformawdcinfouploadAPIRequest) SetTraceInfo(_traceInfo *AwdcInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmalltraceplatformawdcinfouploadAPIRequest) GetTraceInfo() *AwdcInfo {
	return r._traceInfo
}

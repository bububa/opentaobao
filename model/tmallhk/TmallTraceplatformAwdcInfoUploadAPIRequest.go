package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformAwdcInfoUploadAPIRequest
AWDC提交溯源信息 API请求
tmall.traceplatform.awdc.info.upload

天猫溯源-AWDC-上传溯源信息 */
type TmallTraceplatformAwdcInfoUploadAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *AwdcInfo
}

// NewTmallTraceplatformAwdcInfoUploadRequest 初始化TmallTraceplatformAwdcInfoUploadAPIRequest对象
func NewTmallTraceplatformAwdcInfoUploadRequest() *TmallTraceplatformAwdcInfoUploadAPIRequest {
	return &TmallTraceplatformAwdcInfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.awdc.info.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformAwdcInfoUploadAPIRequest) SetTraceInfo(_traceInfo *AwdcInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// Get TraceInfo Getter
func (r TmallTraceplatformAwdcInfoUploadAPIRequest) GetTraceInfo() *AwdcInfo {
	return r._traceInfo
}
